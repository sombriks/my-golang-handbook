package relational

import (
	"0013-databases/model"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // import database drv
	"log"
)

func GetConnection() *sql.DB {
	fmt.Println("get connection for relational database")

	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Insert(c1 *sql.DB, t *model.Todo) int64 {
	return 0
}
