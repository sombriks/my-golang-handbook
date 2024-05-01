package services

import (
	"0016-rest-json/app/configs"
	"0016-rest-json/app/models"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"time"
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
		Where(goqu.C("description").ILike(fmt.Sprint("%", q, "%"))).
		ScanStructs(&result)
	return &result, err
}

func (service *TodoService) Find(id uint64) (*models.Todo, bool, error) {
	var result models.Todo
	found, err := service.dbConfig.DB.From("todos").
		Where(goqu.Ex{"id": id}).ScanStruct(&result)
	return &result, found, err
}

func (service *TodoService) Create(newTodo *models.Todo) (uint64, error) {
	var id uint64
	_, err := service.dbConfig.DB.
		Insert("todos").
		Rows(newTodo).
		Returning("id").
		Executor().ScanVal(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (service *TodoService) Update(id uint64, todo *models.Todo) (int64, error) {
	todo.Id = id
	todo.Updated = time.Now()
	r, err := service.dbConfig.DB.
		Update("todos").
		Where(goqu.Ex{"id": id}).
		Set(todo).Executor().Exec()
	if err != nil {
		return 0, err
	}
	c, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}
	return c, nil
}

func (service *TodoService) Delete(id uint64) (int, error) {

	return 0, nil
}
