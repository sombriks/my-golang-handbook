package model

import "time"

type Todo struct {
	Id          int64
	Description string
	Done        bool
	Created     time.Time
	Updated     time.Time
}

func (t Todo) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = t.Id
	m["description"] = t.Description
	m["done"] = t.Done
	m["created"] = t.Created.Format(time.RFC3339)
	m["updated"] = t.Updated.Format(time.RFC3339)
	return m
}
