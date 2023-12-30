package document_based

import (
	"fmt"
	c "github.com/ostafen/clover"
	"log"
)

func GetConnection() *c.DB {
	fmt.Println("get connection for document-based database")

	db, err := c.Open("clover-db")
	if err != nil {
		log.Panic(err)
	}
	return db
}
