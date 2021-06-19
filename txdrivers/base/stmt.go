package base

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

type stmt struct {
	st *sql.Stmt
}

func (s *stmt) Close() error {
	return s.st.Close()
}

func (s *stmt) NumInput() int {
	return -1
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.st.Exec(mapToInterfaceAny(args)...)
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	rows, err := s.st.Query(mapToInterfaceAny(args)...)
	if err != nil {
		return nil, err
	}
	return buildRows(rows)
}

// ExecContext Implement the "StmtExecContext" interface
func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return s.st.ExecContext(ctx, mapNamedArgs(args)...)
}

// QueryContext Implement the "StmtQueryContext" interface
func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	rows, err := s.st.QueryContext(ctx, mapNamedArgs(args)...)
	if err != nil {
		return nil, err
	}
	return buildRows(rows)
}
