package fixtures

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantPanic   bool
		judgeResult bool
	}{
		{
			name: "unexisted path",
			args: args{
				path: "path/not/existed",
			},
			wantPanic: true,
		},
		{
			name: "existed directory path",
			args: args{
				path: "dummy",
			},
			judgeResult: true,
		},
		{
			name: "existed file path",
			args: args{
				path: "dummy/coupons.yml",
			},
			judgeResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if tt.wantPanic {
					assert.NotNil(t, err)
				}
			}()
			db := OpenDB("mysql", "root:@tcp(127.0.0.1:6606)/?charset=utf8&parseTime=True&loc=Local")
			fixtures := Load(tt.args.path, db)

			if tt.judgeResult {
				count := int(0)
				row := db.QueryRow("SELECT COUNT(*) FROM gofixtures_test.coupons")
				assert.Nil(t, row.Scan(&count))
				assert.Equal(t, 2, count)

				row = db.QueryRow("SELECT COUNT(*) FROM gofixtures_test.users")
				assert.Nil(t, row.Scan(&count))
				assert.Equal(t, 4, count)

				row = db.QueryRow("SELECT COUNT(*) FROM gofixtures_test.admin_users")
				assert.Nil(t, row.Scan(&count))
				assert.Equal(t, 1, count)
				fmt.Printf("%+v\n", fixtures)
			}
		})
	}
}
