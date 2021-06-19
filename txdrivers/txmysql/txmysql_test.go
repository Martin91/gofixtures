package txmysql

import (
	"database/sql"
	"github.com/Martin91/gofixtures/txdrivers/base"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"sync"
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

func TestTxMySQLDriverWithBegin(t *testing.T) {
	tx, err := txMySQLDB.Begin()
	assert.Nil(t, err)
	result, err := tx.Exec("INSERT INTO users (`nickname`) VALUES ('test'), ('value')")
	assert.Nil(t, err)
	rowsAffected, err := result.RowsAffected()
	assert.Equal(t, int64(2), rowsAffected)

	var count int
	row := tx.QueryRow("SELECT COUNT(*) FROM users")
	row.Scan(&count)
	assert.Equal(t, 5, count)

	tx.Rollback()
	row = txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
	row.Scan(&count)
	assert.Equal(t, 3, count)

	tx, _ = txMySQLDB.Begin()
	tx.Exec("INSERT INTO users (`nickname`) VALUES ('july'), ('june')")
	tx.Commit()

	tx, _ = txMySQLDB.Begin()
	tx.Exec("INSERT INTO users (`nickname`) VALUES ('may'), ('april')")
	row = tx.QueryRow("SELECT COUNT(*) FROM users")
	row.Scan(&count)
	assert.Equal(t, 7, count)

	tx.Rollback()
	row = txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
	row.Scan(&count)
	assert.Equal(t, 5, count)

	driver := txMySQLDB.Driver().(base.TxDriverIface)
	driver.ManualRollback()
	row = txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
	err = row.Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 3, count)
}

func TestTxMySQLDriverWithGoroutines(t *testing.T) {
	wg := sync.WaitGroup{}
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		tx, _ := txMySQLDB.Begin()
		result, err := tx.Exec("INSERT INTO users (`nickname`) VALUES ('may'), ('april')")
		assert.Nil(t, err)
		rowsAffected, err := result.RowsAffected()
		assert.Nil(t, err)
		assert.Equal(t, 2, int(rowsAffected))
		ch1 <- true
		select {
		case <- ch2:
			tx.Rollback()
			ch1 <- true
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		var tx *sql.Tx
		select {
		case <- ch1:
			tx, _ = txMySQLDB.Begin()
			tx.Exec("INSERT INTO users (`nickname`) VALUES ('july'), ('Sep')")
			var count int
			row := txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
			row.Scan(&count)
			assert.Equal(t, 7, count) // both goroutines run in a same connection actually, so there is no isolation
			ch2 <- true
		}

		select {
		case <- ch1:
			var count int
			row := txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
			row.Scan(&count)
			assert.Equal(t, 3, count) // the first goroutine will rollback to the earlier savepoint

			tx.Rollback()
			row = txMySQLDB.QueryRow("SELECT COUNT(*) FROM users")
			row.Scan(&count)
			assert.Equal(t, 3, count)
		}
	}()

	wg.Wait()
}