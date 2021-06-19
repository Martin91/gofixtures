// Package txmysql registers a txmysql sql driver
package txpostgresql

import (
	"database/sql"

	"github.com/Martin91/gofixtures/txdrivers/base"
)

type TxMySQLDriver struct {
	base.TxDriver
}

func init() {
	sql.Register("txpostgresql", &TxMySQLDriver{
		TxDriver: base.TxDriver{
			RealDriver: "postgres",
			SavePoint: &base.DefaultSavePoint{},
		},
	})
}
