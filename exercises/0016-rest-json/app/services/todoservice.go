package services

import (
	"0016-rest-json/app/configs"
	"0016-rest-json/app/models"
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

type TodoService struct {
	dbConfig *configs.DBConfig
}

func NewTodoService(dbConfig *configs.DBConfig) *TodoService {
	var service TodoService
	service.dbConfig = dbConfig
	return &service
}

func (service *TodoService) List(q string) (*[]models.Todo, error) {
	var result []models.Todo
	err := service.dbConfig.DB.From("todos").
		Where(goqu.C("description").ILike(fmt.Sprint("%", q,"%"))).
		ScanStructs(&result)
	return &result, err
}

func (service *TodoService) Find(id uint64) (*models.Todo, error) {
	var result models.Todo

	return &result, nil
}

func (service *TodoService) Create(newTodo *models.Todo) (uint64, error) {

	return 0, nil
}

func (service *TodoService) Update(id uint64, todo *models.Todo) (int, error) {

	return 0, nil
}

func (service *TodoService) Delete(id uint64) (int, error) {

	return 0, nil
}
