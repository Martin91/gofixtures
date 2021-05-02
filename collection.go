package fixtures

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// Collection maps to a table
type Collection struct {
	DbName    string              `yaml:"db"`
	TableName string              `yaml:"table_name"`
	Rows      map[string]*Fixture `yaml:"rows"`
}

func (c *Collection) getTableName() string {
	if c.DbName != "" {
		return fmt.Sprintf("`%s`.`%s`", c.DbName, c.TableName)
	}
	return fmt.Sprintf("`%s`", c.TableName)
}

func (c *Collection) insertData(db *sql.DB) error {
	for name, fixture := range c.Rows {
		if name == "DEFAULT" {
			continue
		}
		if err := fixture.insertRow(db, c.getTableName()); err != nil {
			return errors.WithMessagef(err, "create fixture: %s", name)
		}
	}

	return nil
}
