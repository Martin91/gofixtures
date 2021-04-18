// Package fixtures implements a Ruby on Rails style test fixtures suite
package fixtures

import (
	"database/sql"
	"fmt"
	"os"
)

// Fixtures an utility to hold definitions of fixtures
type Fixtures struct {
	path     string             // path to load yaml definition files recursively
	db       *sql.Conn          // db holds an database connection
	fixtures map[string]Fixture // fixtures loaded and parsed definitions
}

func (f *Fixtures) Load() error {
	return nil
}

type CallbackType uint8

const (
	BeforeCreate CallbackType = iota
	AfterCreate
)

type CallbackFunc func(*Fixture) error

type Fixture struct {
	Columns   map[string]interface{}          // Columns fields and values of an single object
	Callbacks map[CallbackType][]CallbackFunc // Callbacks
}

func Load(path string, db *sql.Conn) (*Fixtures, error) {
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
		path: path,
		db:   db,
	}
	return fixtures, fixtures.Load()
}
