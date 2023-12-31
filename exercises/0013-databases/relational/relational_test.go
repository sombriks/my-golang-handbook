package relational

import (
	"0013-databases/model"
	"fmt"
	"testing"
	"time"
)

func TestRelationalOperations(t *testing.T) {

	c1 := GetConnection()
	defer c1.Close()

	newTodo := model.Todo{
		Description: "put trash out",
		Created:     time.Now(),
		Done:        false,
	}

	n1 := Insert(c1, &newTodo)

	if n1 == 0 {
		t.Fail()
	}

	l1 := List(c1, "")

	if l1 == nil {
		t.Fail()
	}

	if len(l1) > 0 {
		l1[0].Description = "Updated description"
		l1[0].Done = true
		Update(c1, l1[0])
	}

	if len(l1) > 0 {
		todo := Find(c1, l1[0].Id)
		fmt.Printf("todo for id %d: %+v\n", l1[0].Id, todo)
	}

	if len(l1) > 0 {
		Del(c1, l1[0].Id)
	}
}
