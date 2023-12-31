package main

import (
	doc "0013-databases/document-based"
	kv "0013-databases/key-value-based"
	"0013-databases/model"
	rel "0013-databases/relational"
	"fmt"
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

	newTodo := model.Todo{
		Description: "put trash out",
		Created:     time.Now(),
		Done:        false,
	}

	// insert todo
	n1 := rel.Insert(c1, &newTodo)
	n2 := doc.Insert(c2, &newTodo)
	n3 := kv.Insert(c3, &newTodo)

	fmt.Printf("todo inserted\nrelational key: %d\ndocument key: %s\nkey/value key: %d\n", n1, n2, n3)

	// list todos's
	l1 := rel.List(c1, "")
	for _, e := range l1 {
		fmt.Printf("%+v\n", e)
	}
	l2 := doc.List(c2, "")
	for _, e := range l2 {
		fmt.Printf("%+v\n", e)
	}
	l3 := kv.List(c3, "2023") // We query keys on this guy
	for _, e := range l3 {
		fmt.Printf("%+v\n", e)
	}

	// update todo

	// find todo

	// delete todo

}
