package document_based

import (
	"0013-databases/model"
	"fmt"
	c "github.com/ostafen/clover"
	"log"
)

func GetConnection() *c.DB {
	fmt.Println("get connection for document-based database")

	db, err := c.Open("todo-doc")
	if err != nil {
		log.Panic(err)
	}

	_ = db.CreateCollection("todos")
	return db
}

func Insert(db *c.DB, todo *model.Todo) string {
	doc := c.NewDocumentOf(todo.ToMap())
	id, _ := db.InsertOne("todos", doc)
	return id
}
