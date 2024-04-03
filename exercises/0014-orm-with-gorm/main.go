package main

import (
	"0014-orm-with-gorm/app"
	"log"
)

// Some operations
func main() {
	example14 := app.NewExample14()
	example14.Insert("Alice", "Down Way Street, 101")
	example14.Insert("Bob", "Baker Avenue, 10")
	example14.Insert("James", "101 Street, 1500")
	example14.Insert("Jessica", "Margaret Boulevard, 123")
	result, _ := example14.ListByName("Alice")
	contact := &(*result)[0]
	contact.Name = "Mary"
	example14.Update(contact)
	contact, _ = example14.Find(contact.ID)
	example14.Delete(contact.ID)
	_, err := example14.Find(contact.ID)
	log.Printf("%s\n", err)
	example14.CloseDb()
}
