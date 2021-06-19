// Package base defines common logic among different transational drivers
package base

import (
	"database/sql"
	"database/sql/driver"
	"sync"
)

type TxDriverIface interface {
	ManualRollback()
	DeleteConn(string)
}

// TxDriver root type for all other inheriting and transactional drivers
type TxDriver struct {
	sync.Mutex
	SavePoint   SavePointIface
	RealDriver  string           // underlying sql driver used to establish real db connections
	connections map[string]*conn // maintains connections isolated by different DSNs
}

// ManualRollback allow user to control the rollback point
func (d *TxDriver) ManualRollback() {
	d.Lock()
	defer d.Unlock()

	for _, connection := range d.connections {
		connection.ManualRollback()
	}
}

// Open ensures single transaction by isolating connections by their dsn
func (d *TxDriver) Open(dsn string) (driver.Conn, error) {
	d.Lock()
	defer d.Unlock()

	connection, existed := d.connections[dsn]
	if !existed {
		if db, err := sql.Open(d.RealDriver, dsn); err != nil {
			return nil, err
		} else {
			connection = &conn{
				driver:        d,
				db:            db,
				dsn:           dsn,
				savePointImpl: d.SavePoint,
			}
			if d.connections == nil {
				d.connections = map[string]*conn{}
			}
			d.connections[dsn] = connection
		}
	}
	connection.occupied++
	return connection, nil
}

func (d *TxDriver) DeleteConn(dsn string) {
	d.Lock()
	defer d.Unlock()

	delete(d.connections, dsn)
}
