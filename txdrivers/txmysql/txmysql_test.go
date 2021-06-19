package txmysql

import (
	"database/sql"
	"github.com/Martin91/gofixtures/txdrivers/base"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var (
	mysqlDB   *sql.DB
	txMySQLDB *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	mysqlDB, err = sql.Open("mysql", "root@(127.0.0.1:6606)/gofixtures_test")
	if err != nil {
		log.Fatalf("failed to open mysql db, err: %s", err)
	}
	txMySQLDB, err = sql.Open("txmysql", "root@(127.0.0.1:6606)/gofixtures_test")
	if err != nil {
		log.Fatalf("failed to open txmysql db, err: %s", err)
	}
	os.Exit(m.Run())
}

func TestTxMySQLDriverRunsInAGlobalTransaction(t *testing.T) {
	count := 0
	row := mysqlDB.QueryRow("SELECT COUNT(*) FROM users")
	err := row.Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 3, count)

	row = txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
	err = row.Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 3, count)

	result, err := txMySQLDB.Exec("INSERT INTO users (`nickname`) VALUES ('hello'), ('world')")
	assert.Nil(t, err)
	rowsAffected, err := result.RowsAffected()
	assert.Nil(t, err)
	assert.Equal(t, int64(2), rowsAffected)

	row = mysqlDB.QueryRow("SELECT COUNT(*) FROM users")
	err = row.Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 3, count)

	row = txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
	err = row.Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 5, count)

	driver := txMySQLDB.Driver().(base.TxDriverIface)
	driver.ManualRollback()
	row = txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
	err = row.Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 3, count)
}
