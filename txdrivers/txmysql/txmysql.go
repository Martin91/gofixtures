// Package txmysql registers a txmysql sql driver
package txmysql

import (
	"database/sql"

	"github.com/Martin91/gofixtures/txdrivers/base"
)

type TxMySQLDriver struct {
	base.TxDriver
}

func init() {
	sql.Register("txmysql", &TxMySQLDriver{
		TxDriver: base.TxDriver{
			RealDriver: "mysql",
		},
	})
}
