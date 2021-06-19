// Package txmysql registers a txmysql sql driver
package txmysql

import (
	"database/sql"

	"github.com/Martin91/gofixtures/txdrivers/base"
)

type driver struct {
	base.TxDriver
}

func init() {
	sql.Register("txmysql", &driver{
		TxDriver: base.TxDriver{
			RealDriver: "mysql",
			SavePoint: &base.DefaultSavePoint{},
		},
	})
}
