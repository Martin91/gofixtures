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
		wantErr     bool
		judgeResult bool
	}{
		{
			name: "unexisted path",
			args: args{
				path: "path/not/existed",
			},
			wantErr: true,
		},
		{
			name: "existed directory path",
			args: args{
				path: "dummy",
			},
			judgeResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := OpenDB("mysql", "root:@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local")
			fixtures, err := Load(tt.args.path, db)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.judgeResult {
				assert.NotEmpty(t, fixtures.collections["coupons"])
				assert.NotEmpty(t, fixtures.collections["users"])
				assert.NotEmpty(t, fixtures.collections["administrators"])
				fmt.Printf("%+v\n", fixtures)
			}
		})
	}
}
