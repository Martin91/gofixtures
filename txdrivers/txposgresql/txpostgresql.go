// Package txpostgresql registers a txpostgresql sql driver
package txpostgresql

import (
	"database/sql"

	"github.com/Martin91/gofixtures/txdrivers/base"
)

type driver struct {
	base.TxDriver
}

func init() {
	sql.Register("txpostgresql", &driver{
		TxDriver: base.TxDriver{
			RealDriver: "postgres",
			SavePoint: &base.DefaultSavePoint{},
		},
	})
}
