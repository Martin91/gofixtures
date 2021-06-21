package fixtures

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
)

var evaluatorPattern = regexp.MustCompile(`^\${(?P<evaluator>[a-zA-Z0-9]+)\(\)}$`)

type Fixture struct {
	Columns map[string]interface{} // Columns fields and values of an single object
}

func (f *Fixture) parseValue(value interface{}) interface{} {
	str, ok := value.(string)
	if !ok {
		return value
	}

	if !strings.HasPrefix(str, "${") || !strings.HasSuffix(str, "}") {
		return value
	}
	result := evaluatorPattern.FindStringSubmatch(str)
	if result == nil {
		return value
	}

	evaluator, ok := evaluators[result[1]]
	if !ok {
		return value
	}

	return evaluator()
}

func (f *Fixture) insertRow(db *sql.DB, tableName string) error {
	sqlPattern := "INSERT INTO %s (%s) VALUES (%s)"
	columns := make([]string, 0, len(f.Columns))
	values := make([]interface{}, 0, len(f.Columns))
	placeholders := make([]string, 0, len(f.Columns))

	for name, value := range f.Columns {
		columns = append(columns, fmt.Sprintf("`%s`", name))
		values = append(values, f.parseValue(value))
		placeholders = append(placeholders, "?")
	}

	columnsSQL := strings.Join(columns, ", ")
	valuesSQL := strings.Join(placeholders, ", ")
	sql := fmt.Sprintf(sqlPattern, tableName, columnsSQL, valuesSQL)
	_, err := db.Exec(sql, values...)
	return err
}
