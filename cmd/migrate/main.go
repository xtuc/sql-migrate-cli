package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	user = os.Getenv("MYSQL_USER")
	pass = os.Getenv("MYSQL_PASS")
	host = os.Getenv("MYSQL_HOST")
	db   = os.Getenv("MYSQL_DB")
)

func connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pass, host, db)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	migrations := &migrate.FileMigrationSource{
		Dir: "mysql-migrations",
	}

	db, err := connect()

	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
