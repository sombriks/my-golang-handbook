package relational

import (
	"0013-databases/model"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // import database drv
	"log"
	"time"
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
		result = append(result, model.FromRows(rows))
	}
	return result
}

func Update(db *sql.DB, todo model.Todo) int64 {
	todo.Updated = time.Now()
	result, err := db.Exec(`
		update todo set 
		description = ?,
		done = ?,
		updated = ?
		where id = ?
	`, todo.Description, todo.Done, todo.Updated, todo.Id)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := result.RowsAffected()
	return affected
}

func Find(db *sql.DB, id int64) model.Todo {
	row := db.QueryRow("select * from todo where id = ?", id)
	return model.FromRow(row)
}

func Del(db *sql.DB, id int64) int64 {
	result, err := db.Exec("delete from todo where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := result.RowsAffected()
	return affected
}
