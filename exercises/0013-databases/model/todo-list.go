package model

import (
	"database/sql"
	c "github.com/ostafen/clover"
	"time"
)

type Todo struct {
	Id          int64
	Description string
	Done        bool
	Created     time.Time
	Updated     time.Time
}

func (t Todo) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = t.Id
	m["description"] = t.Description
	m["done"] = t.Done
	m["created"] = t.Created.Format(time.RFC3339)
	m["updated"] = t.Updated.Format(time.RFC3339)
	return m
}

func FromRows(row *sql.Rows) Todo {
	var todo Todo
	_ = row.Scan(&todo.Id, &todo.Description, &todo.Done, &todo.Created, &todo.Updated)
	return todo
}

func FromRow(row *sql.Row) Todo {
	var todo Todo
	_ = row.Scan(&todo.Id, &todo.Description, &todo.Done, &todo.Created, &todo.Updated)
	return todo
}

func FromDoc(doc *c.Document) Todo {
	var todo Todo
	todo.Id = doc.Get("id").(int64)                    // Cannot convert an expression of the type 'interface{}' to the type 'int64'
	todo.Description = doc.Get("description").(string) // https://go.dev/tour/methods/15
	todo.Done = doc.Get("done").(bool)
	todo.Created, _ = time.Parse(time.RFC3339, doc.Get("created").(string))
	todo.Updated, _ = time.Parse(time.RFC3339, doc.Get("updated").(string))
	return todo
}
