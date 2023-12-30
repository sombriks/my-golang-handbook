package main

import (
	doc "0013-databases/document-based"
	kv "0013-databases/key-value-based"
	rel "0013-databases/relational"
	"time"
)

type Todo struct {
	Id          int64
	description string
	done        bool
	created     time.Time
}

func main() {

	// recover db connections
	c1 := rel.GetConnection()
	c2 := doc.GetConnection()
	c3 := kv.GetConnection()

	// defer connection close
	defer c1.Close()
	defer c2.Close()
	defer c3.Close()

	// create todo

	// update todo

	// list todos's

	// find todo

	// delete todo

}
