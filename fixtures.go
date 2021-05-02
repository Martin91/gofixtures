// Package fixtures implements a Ruby on Rails style test fixtures suite
package fixtures

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/DATA-DOG/go-txdb"
	"github.com/pkg/errors"
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

// registeredDrivers Ensure that registering each driver exactly once globally
var registeredDrivers = sync.Map{}

// OpenDB setup a transactional db to automatically rollback db changes
func OpenDB(driverName, dsn string) (*sql.DB, error) {
	var txDriverName string
	txDriver, ok := registeredDrivers.Load(driverName)
	if !ok {
		txDriverName = fmt.Sprintf("tx%sdb", driverName)
		txdb.Register(txDriverName, driverName, dsn)
		registeredDrivers.Store(driverName, txDriverName)
	} else {
		txDriverName = txDriver.(string)
	}
	return sql.Open(txDriverName, dsn)
}

// Load parse yaml files under the directory specified by `path`
func Load(path string, db *sql.DB) (*Fixtures, error) {
	stat, err := os.Stat(path)
	if err == nil {
		if !stat.IsDir() {
			err = fmt.Errorf("path %s is not a directory", path)
		}
	}
	if err != nil {
		return nil, err
	}

	fixtures := &Fixtures{
		path:        path,
		db:          db,
		collections: map[string]*Collection{},
	}
	return fixtures, fixtures.Load()
}
