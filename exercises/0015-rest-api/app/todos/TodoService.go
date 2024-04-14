package todos

import (
	"0015-rest-api/app/configs"
	"fmt"
)

type TodoService struct {
	database *configs.Database
}

func NewTodoService(database *configs.Database) *TodoService {
	var service TodoService
	service.database = database
	return &service
}

func (service *TodoService) List(q string) (*[]TodoItem, error) {
	var todos []TodoItem
	like := fmt.Sprintf("%%%s%%", q)
	err := service.database.Db.
		Where("lower(description) like lower(?)", like).
		Find(&todos).Error
	return &todos, err
}

func (service *TodoService) Find(id uint) (*TodoItem, error) {
	var todo TodoItem
	err := service.database.Db.First(&todo, id).Error
	return &todo, err
}

func (service *TodoService) Insert(item TodoItem) (uint64, error) {
	err := service.database.Db.Create(&item).Error
	return item.Id, err
}

func (service *TodoService) Update(id uint64, todo TodoItem) (int64, error) {
	todo.Id = id
	result := service.database.Db.Save(&todo)
	return result.RowsAffected, result.Error
}

func (service *TodoService) Delete(id uint64) (int64, error) {
	result := service.database.Db.Delete(&TodoItem{}, id)
	return result.RowsAffected, result.Error
}
