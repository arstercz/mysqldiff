/*config read to verify normal user*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func dbh(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return db, err
	}
	return db, nil
}

func Query(db *sql.DB, q string) (*sql.Rows, error) {
	return db.Query(q)
}

func QueryRow(db *sql.DB, q string) *sql.Row {
	return db.QueryRow(q)
}

func ExecQuery(db *sql.DB, q string) (sql.Result, error) {
	return db.Exec(q)
}

func insertlog(db *sql.DB, t *mysqlParams) bool {
	insertSql := `
	insert into mysql_diff(host, port, db, tag, changes, create_time) values('%s', %d, '%s', '%s', '%s', now())
	`
	_, err := ExecQuery(db, fmt.Sprintf(insertSql, t.host, t.port, t.db, t.tag, sql_escape(t.changes)))
	if err != nil {
		return false
	}
	return true
}
