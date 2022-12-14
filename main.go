package main

import (
	"Slices/database"
	"context"
	"database/sql"
	"fmt"
)

func main() {
	var id int
	var err error
	ctx := context.Background()

	//create connection
	db := database.CreateConnection(ctx)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// test query
	err = db.QueryRowContext(ctx, "SELECT id from book").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("id")
	fmt.Println("All good")

}
