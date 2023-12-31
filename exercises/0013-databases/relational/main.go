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

	init := `
		create table if not exists todo(
		    id integer primary key autoincrement, 
		    description text, 
		    done boolean, 
		    created timestamp, 
		    updated timestamp
		);
	`

	_, _ = db.Exec(init)

	return db
}

func Insert(db *sql.DB, todo *model.Todo) int64 {

	result, _ := db.Exec(`
		insert into todo (description,done,created,updated)
		values(?,?,?,?)
	`, todo.Description, todo.Done, todo.Created, todo.Updated)

	id, _ := result.LastInsertId()

	return id
}

func List(db *sql.DB, q string) []model.Todo {
	rows, err := db.Query(`
		select * from todo
	 	where lower(description) 
	  	like lower('%'||?||'%')`, q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	result := make([]model.Todo, 0)
	for rows.Next() {
		result = append(result, model.FromRow(rows))
	}
	return result
}
