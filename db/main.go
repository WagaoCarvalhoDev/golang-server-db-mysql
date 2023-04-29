package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:passdb@tcp(localhost:3306)/serverdb")
	if err != nil {
		panic(err)
	}

	//defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS serverdb")
	exec(db, "USE serverdb")
	//exec(db, "DROP TABLE IF EXISTS users")
	exec(db, `CREATE TABLE IF NOT EXISTS users(
		id INTEGER AUTO_INCREMENT,
		name VARCHAR(80),
		PRIMARY KEY(id)
	)`)

	return db, nil
}
