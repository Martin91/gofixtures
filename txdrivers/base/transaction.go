package base

type tx struct {
	id   string
	conn *conn
}

// Commit release a savepoint if there is one, otherwise do nothing
func (tx *tx) Commit() error {
	if tx.conn.savePointImpl == nil {
		return nil // save point is not supported
	}

	tx.conn.Lock()
	defer tx.conn.Unlock()

	connTx, err := tx.conn.beginOnce()
	if err != nil {
		return err
	}

	_, err = connTx.Exec(tx.conn.savePointImpl.Release(tx.id))
	return err
}

// Rollback discard a savepoint if there is one, otherwise do nothing
func (tx *tx) Rollback() error {
	if tx.conn.savePointImpl == nil {
		return nil // save point is not supported
	}

	tx.conn.Lock()
	defer tx.conn.Unlock()

	connTx, err := tx.conn.beginOnce()
	if err != nil {
		return err
	}

	_, err = connTx.Exec(tx.conn.savePointImpl.Rollback(tx.id))
	return err
}
