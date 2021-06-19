package base

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"sync"
)

type conn struct {
	sync.Mutex
	driver        TxDriverIface
	dsn           string
	occupied      int
	tx            *sql.Tx
	db            *sql.DB
	savePoints    int
	savePointImpl SavePointIface
}

// ManualRollback force to rollback the underlying transaction
func (c *conn) ManualRollback() error {
	c.Lock()
	defer c.Unlock()

	if c.tx != nil {
		if err := c.tx.Rollback(); err != nil {
			return err
		}
		c.tx = nil
	}
	return nil
}

func (c *conn) beginOnce() (*sql.Tx, error) {
	// A normal driver.Conn is assured that it is used by only one goroutine,
	// but we break this convention in TxDriver.Open method so we need to gurantee
	// that this method run simultaneously
	c.Lock()
	defer c.Unlock()

	if c.tx == nil {
		tx, err := c.db.Begin()
		if err != nil {
			return nil, err
		}
		c.tx = tx
	}
	return c.tx, nil
}

// Prepare prepares a statement
func (c *conn) Prepare(query string) (driver.Stmt, error) {
	tx, err := c.beginOnce()
	if err != nil {
		return nil, err
	}

	st, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &stmt{st: st}, nil
}

func (c *conn) Close() error {
	c.Lock()
	defer c.Unlock()

	c.occupied--
	if c.occupied == 0 {
		if c.tx != nil {
			c.tx.Rollback()
			c.tx = nil
		}
		if err := c.db.Close(); err != nil {
			return err
		}
		c.driver.DeleteConn(c.dsn)
	}
	return nil
}

func (c *conn) Begin() (driver.Tx, error) {
	if c.savePointImpl == nil {
		return &tx{"_", c}, nil // save point is not supported
	}

	c.Lock()
	defer c.Unlock()

	connTx, err := c.beginOnce()
	if err != nil {
		return nil, err
	}

	c.savePoints++
	id := fmt.Sprintf("tx_%d", c.savePoints)
	_, err = connTx.Exec(c.savePointImpl.Create(id))
	if err != nil {
		return nil, err
	}
	return &tx{id, c}, nil
}

func (c *conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	c.Lock()
	defer c.Unlock()

	tx, err := c.beginOnce()
	if err != nil {
		return nil, err
	}

	// query rows
	rs, err := tx.Query(query, mapToInterfaceAny(args)...)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	return buildRows(rs)
}

func (c *conn) beginTxOnce(ctx context.Context) (*sql.Tx, error) {
	if c.tx == nil {
		tx, err := c.db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return nil, err
		}
		c.tx = tx
	}
	return c.tx, nil
}

// Implement the "QueryerContext" interface
func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	c.Lock()
	defer c.Unlock()

	tx, err := c.beginTxOnce(ctx)
	if err != nil {
		return nil, err
	}

	rs, err := tx.QueryContext(ctx, query, mapNamedArgs(args)...)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	return buildRows(rs)
}

// Implement the "ExecerContext" interface
func (c *conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	c.Lock()
	defer c.Unlock()

	tx, err := c.beginTxOnce(ctx)
	if err != nil {
		return nil, err
	}

	return tx.ExecContext(ctx, query, mapNamedArgs(args)...)
}

// Implement the "ConnBeginTx" interface
func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return c.Begin()
}

// Implement the "ConnPrepareContext" interface
func (c *conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	c.Lock()
	defer c.Unlock()

	tx, err := c.beginTxOnce(ctx)
	if err != nil {
		return nil, err
	}

	st, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &stmt{st: st}, nil
}

// Implement the "Pinger" interface
func (c *conn) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}
