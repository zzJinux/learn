package learnstd

import (
	"database/sql"
	"fmt"
	"testing"

	"learn/database"
	. "learn/database/u"

	_ "github.com/go-sql-driver/mysql"
	rq "github.com/stretchr/testify/require"
)

// connect creates a connection to the MySQL database.
func connect(port uint16) *sql.DB {
	// DSN format: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	// parseTime=true changes the output type of DATE and DATETIME values to time.Time instead of []byte / string.
	// The date or datetime like 0000-00-00 00:00:00 is converted into zero value of time.Time.
	// https://github.com/go-sql-driver/mysql#parsetime
	dbname := "learn_sandbox"
	dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:%d)/?parseTime=true&multiStatements=true", port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + "; USE " + dbname + ";"); err != nil {
		panic(err)
	}
	return db
}

type Model struct {
	ID  string
	Val string
}

func resetTable(db *sql.DB, tableName string) {
	_, err := db.Exec("DROP TABLE IF EXISTS " + tableName)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE %s (id VARCHAR(32), val VARCHAR(32))", tableName))
	if err != nil {
		panic(err)
	}
}

func TestExample(t *testing.T) {
	db := connect(database.EnvPort)
	resetTable(db, "models")

	Must1(db.Exec("INSERT INTO models VALUES(?,?)", "a", "court"))
	Must1(db.Exec("INSERT INTO models VALUES(?,?)", "b", "bamboo"))

	// Query a single row.
	{
		var id, val string
		var row *sql.Row

		row = db.QueryRow("SELECT * FROM models WHERE id = ?", "a")
		Must(row.Scan(&id, &val))
		rq.Equal(t, "a", id)
		rq.Equal(t, "court", val)

		row = db.QueryRow("SELECT * FROM models WHERE id = ?", "b")
		Must(row.Scan(&id, &val))
		rq.Equal(t, "b", id)
		rq.Equal(t, "bamboo", val)
	}

	// Query multiple rows.
	{
		rows := Must1(db.Query("SELECT * FROM models WHERE CHAR_LENGTH(id) > 0"))
		rr := make([]Model, 2)
		for i := 0; rows.Next(); i++ {
			Must(rows.Scan(&rr[i].ID, &rr[i].Val))
		}
		rq.ElementsMatch(t, []Model{{"a", "court"}, {"b", "bamboo"}}, rr)
	}

	// Transaction
	{
		tx := Must1(db.Begin())
		Must1(tx.Exec("INSERT INTO models VALUES(?,?)", "c", "gugugu"))
		Must(tx.Rollback())
		rq.ErrorIs(t, tx.Rollback(), sql.ErrTxDone)
	}

	// Query multiple rows (verify rollback)
	{
		rows := Must1(db.Query("SELECT * FROM models"))
		rr := make([]Model, 0, 3)
		for i := 0; rows.Next(); i++ {
			var r Model
			Must(rows.Scan(&r.ID, &r.Val))
			rr = append(rr, r)
		}
		rq.Len(t, rr, 2)
	}
}
