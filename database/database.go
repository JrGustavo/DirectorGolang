package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func CreateConnection(ctx context.Context) *sql.DB {
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		panic("DB_HOST no set")
	}
	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		panic("DB_USER not set")
	}

	dbname, ok = os.LookupEnv("DB_NAME")
	if !ok {
		panic("DB_NAME not set")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		panic("DB_PORT not set")
	}
	connectionString := fmt.Sprintf("%s:1234@tcp(%s:%d)/%s", user, host, dbname)

	db, err := sql.Open("mysql, connectionString", "root:1234@tcp(localhost:3306)/books")
	if err != nil {
		panic(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	return db
}
