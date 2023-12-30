package key_value_based

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func GetConnection() *leveldb.DB {
	fmt.Println("get connection for key/value database")
	db, err := leveldb.OpenFile("todo-kv", nil)
	if err != nil {
		log.Panic(err)
	}
	return db
}
