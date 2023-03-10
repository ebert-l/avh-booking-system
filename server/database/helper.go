package database

import (
	"database/sql"
	"log"
	"os"
)

// HandleDatabaseError logs error to console
func HandleDatabaseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// HandleTxError rolls back transaction and logs error to console
func HandleTxError(tx *sql.Tx, err error) {
	// TODO make safer (copies lock value)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}
}

// TxRowsAffected handles transaction error and prints affected rows to console.
func TxRowsAffected(res sql.Result, tx *sql.Tx) {
	lastID, err := res.LastInsertId()
	HandleTxError(tx, err)
	rowCnt, err := res.RowsAffected()
	HandleDatabaseError(err)
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}

// Retrieve Environment variables and use a fallback if it is not yet set
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
