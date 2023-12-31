package key_value_based

import (
	"0013-databases/model"
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"time"
)

func GetConnection() *leveldb.DB {
	fmt.Println("get connection for key/value database")
	db, err := leveldb.OpenFile("todo-kv", nil)
	if err != nil {
		log.Panic(err)
	}
	return db
}

func Insert(db *leveldb.DB, todo *model.Todo) int64 {
	todo.Id = todo.Created.Unix()
	_todo, _ := json.Marshal(todo)
	_ = db.Put([]byte(todo.Created.Format(time.RFC3339)), _todo, nil)
	return todo.Id
}
