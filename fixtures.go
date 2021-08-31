// Package fixtures implements a Ruby on Rails style test fixtures suite
package fixtures

import (
	"database/sql"
	"fmt"
	"github.com/Martin91/gofixtures/txdrivers"
	"github.com/Martin91/gofixtures/txdrivers/base"
	_ "github.com/Martin91/gofixtures/txdrivers/txmysql"
	_ "github.com/Martin91/gofixtures/txdrivers/txposgresql"
	"github.com/pkg/errors"
	"os"
)

// Fixtures an utility to hold definitions of fixtures
type Fixtures struct {
	path        string                 // path to load yaml definition files recursively
	db          *sql.DB                // db holds an database connection
	collections map[string]*Collection // fixtures loaded and parsed definitions
}

// insertData insert rows
func (f *Fixtures) insertData() error {
	for name, collection := range f.collections {
		if err := collection.insertData(f.db); err != nil {
			return errors.WithMessagef(err, "create collection: %s", name)
		}
	}
	return nil
}

// OpenDB setup a transactional db to automatically rollback db changes, it may panics if any error encountered
// 	driverName is the actually underlying driver, for example, `mysql`
func OpenDB(driverName, dsn string) *sql.DB {
	txDriverName := txdrivers.TxDriverName(driverName)
	if db, err := sql.Open(txDriverName, dsn); err != nil {
		panic(err)
	} else {
		return db
	}
}

// Load parse yaml files under the directory specified by `path`, it may panics if any error encountered
func Load(path string, db *sql.DB) *Fixtures {
	_, err := os.Stat(path)
	var fixtures *Fixtures
	if err == nil {
		fixtures = &Fixtures{
			path:        path,
			db:          db,
			collections: map[string]*Collection{},
		}
		err = fixtures.Load()
	}

	if err != nil {
		panic(err)
	}
	return fixtures
}

// Rollback manually rollback all changes since last call on Load
func Rollback(db *sql.DB) error {
	if txDriver, ok := db.Driver().(base.TxDriverIface); ok {
		txDriver.ManualRollback()
		return nil
	}
	return fmt.Errorf("it seems like that this db is not driven by our transational driver")
}
