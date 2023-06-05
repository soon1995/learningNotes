package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // support MySQL
	_ "github.com/lib/pq"              // support Postgres
)

func main() {
  // ... 
	db, err = sql.Open("postgres", dbname) // ok
	db, err = sql.Open("mysql", dbname)    // ok
	db, err = sql.Open("sqlite3", dbname)  // returns error: unknown driver "sqlite3"
}
