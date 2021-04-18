// Package fixtures implements a Ruby on Rails style test fixtures suite
package fixtures

import (
	"database/sql"
	"fmt"
	"os"
)

// Fixtures an utility to hold definitions of fixtures
type Fixtures struct {
	path        string                 // path to load yaml definition files recursively
	db          *sql.Conn              // db holds an database connection
	collections map[string]*Collection // fixtures loaded and parsed definitions
}

// Collection maps to a table
type Collection struct {
	DbName    string              `yaml:"db"`
	TableName string              `yaml:"table_name"`
	Rows      map[string]*Fixture `yaml:"rows"`
}

type CallbackType uint8

const (
	BeforeCreate CallbackType = iota
	AfterCreate
)

type CallbackFunc func(*Fixture) error

type Fixture struct {
	Columns map[string]interface{} // Columns fields and values of an single object
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
		path:        path,
		db:          db,
		collections: map[string]*Collection{},
	}
	return fixtures, fixtures.Load()
}
