package base

import (
	"database/sql"
	"database/sql/driver"
)

func mapToInterfaceAny(args ...interface{}) []interface{} {
	return args
}

func buildRows(r *sql.Rows) (driver.Rows, error) {
	set := &rowSets{}
	rs := &rows{}
	if err := rs.read(r); err != nil {
		return set, err
	}
	set.sets = append(set.sets, rs)
	for r.NextResultSet() {
		rss := &rows{}
		if err := rss.read(r); err != nil {
			return set, err
		}
		set.sets = append(set.sets, rss)
	}
	return set, nil
}

func mapNamedArgs(args []driver.NamedValue) (res []interface{}) {
	res = make([]interface{}, len(args))
	for i := range args {
		name := args[i].Name
		if name != "" {
			res[i] = sql.Named(name, args[i].Value)
		} else {
			res[i] = args[i].Value
		}
	}
	return
}
