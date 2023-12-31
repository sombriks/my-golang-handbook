package document_based

import (
	"0013-databases/model"
	"fmt"
	"testing"
	"time"
)

func TestDocumentOperations(t *testing.T) {

	c2 := GetConnection()
	defer c2.Close()

	newTodo := model.Todo{
		Description: "put trash out",
		Created:     time.Now(),
		Done:        false,
	}

	Insert(c2, &newTodo)

	l2 := List(c2, "")
	for _, e := range l2 {
		fmt.Printf("%+v\n", e)
	}

	if len(l2) > 0 {
		l2[0].Description = "Updated description"
		l2[0].Done = true
		Update(c2, l2[0])
	}

	if len(l2) > 0 {
		todo := Find(c2, l2[0].Id)
		fmt.Printf("todo for id %d: %+v\n", l2[0].Id, todo)
	}

	if len(l2) > 0 {
		Del(c2, l2[0].Id)
	}
}
