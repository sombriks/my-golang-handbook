package services

import "0016-rest-json/app/configs"

type TodoService struct {
	dbConfig *configs.DBConfig
}

func NewTodoService(dbConfig *configs.DBConfig) *TodoService {
	var service TodoService
	service.dbConfig = dbConfig
	return &service
}
