package fixtures

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	type args struct {
		path string
		db   *sql.Conn
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
			fixtures, err := Load(tt.args.path, tt.args.db)
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
