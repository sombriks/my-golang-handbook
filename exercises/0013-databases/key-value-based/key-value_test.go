package key_value_based

import (
	"0013-databases/model"
	"fmt"
	"testing"
	"time"
)

func TestKeyValueOperations(t *testing.T) {

	c3 := GetConnection()
	defer c3.Close()

	newTodo := model.Todo{
		Description: "put trash out",
		Created:     time.Now(),
		Done:        false,
	}
	n3 := Insert(c3, &newTodo)

	if n3 == 0 {
		t.Fail()
	}

	l3 := List(c3, "2023") // We query keys on this guy
	for _, e := range l3 {
		fmt.Printf("%+v\n", e)
	}

	if len(l3) > 0 {
		l3[0].Description = "Updated description"
		l3[0].Done = true
		Update(c3, l3[0])
	}

	if len(l3) > 0 {
		todo := Find(c3, l3[0].Id)
		fmt.Printf("todo for id %d: %+v\n", l3[0].Id, todo)
	}

	if len(l3) > 0 {
		Del(c3, l3[0].Id)
	}
}
