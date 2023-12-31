package document_based

import (
	"0013-databases/model"
	"fmt"
	c "github.com/ostafen/clover"
	"log"
	"time"
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
	todo.Id = todo.Created.Unix() // It's an insert
	doc := c.NewDocumentOf(todo.ToMap())
	id, _ := db.InsertOne("todos", doc)
	return id
}

func List(db *c.DB, q string) []model.Todo {
	docs, err := db.Query("todos").Where(c.Field("description").Like(q)).FindAll()
	if err != nil {
		log.Panic(err)
	}
	result := make([]model.Todo, 0)
	for _, doc := range docs {
		result = append(result, model.FromDoc(doc))
	}
	return result
}

func Update(db *c.DB, todo model.Todo) {
	todo.Updated = time.Now()

	err := db.
		Query("todos").
		Where(c.Field("id").
			Eq(todo.Id)).
		Update(todo.ToMap())

	if err != nil {
		log.Panic(err)
	}
}
