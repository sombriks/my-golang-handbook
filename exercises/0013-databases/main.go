package main

import (
	doc "0013-databases/document-based"
	kv "0013-databases/key-value-based"
	"0013-databases/model"
	rel "0013-databases/relational"
	"time"
)

func main() {

	// recover db connections
	c1 := rel.GetConnection()
	c2 := doc.GetConnection()
	c3 := kv.GetConnection()

	// defer close
	defer c1.Close()
	defer c2.Close()
	defer c3.Close()

	// create todo
	newTodo := model.Todo{
		Description: "put trash out",
		Created:     time.Now(),
		Done:        false,
	}

	n1 := rel.Insert(c1, &newTodo)
	n2 := doc.Insert(c2, &newTodo)
	n3 := kv.Insert(c3, &newTodo)

	// update todo

	// list todos's

	// find todo

	// delete todo

}
