package service

import (
	"0014-orm-with-gorm/app/config"
	"0014-orm-with-gorm/app/model"
	"errors"
	"fmt"
)

type Agenda struct {
	database *config.Database
}

func NewAgenda(database *config.Database) *Agenda {
	var agenda Agenda
	agenda.database = database
	return &agenda
}

func (agenda *Agenda) ListByName(q string) (*[]model.Contact, error) {
	// https://gorm.io/docs/preload.html
	var result []model.Contact
	like := fmt.Sprintf("%%%s%%", q)
	err := agenda.database.Db.
		Preload("Addresses").
		Where("lower(name) like lower(?)", like).
		Find(&result).Error
	return &result, err
}

func (agenda *Agenda) ListByAddress(q string) (*[]model.Contact, error) {
	var result []model.Contact
	like := fmt.Sprintf("%%%s%%", q)
	err := agenda.database.Db.
		Preload("Addresses").
		Where(`id in (
				select contact_id 
				from addresses 
				where lower(street) like lower(?))`, like).
		Find(&result).Error
	return &result, err
}

func (agenda *Agenda) Find(id uint) (*model.Contact, error) {
	var result model.Contact
	err := agenda.database.Db.First(&result, id).Error
	return &result, err
}

func (agenda *Agenda) Insert(contact *model.Contact) (uint, error) {
	err := agenda.database.Db.Create(contact).Error
	return contact.ID, err
}

func (agenda *Agenda) Update(contact *model.Contact) error {
	if contact.ID == 0 {
		return errors.New("missing ID")
	}
	return agenda.database.Db.Save(contact).Error
}

func (agenda *Agenda) Delete(id uint) error {
	return agenda.database.Db.Delete(&model.Contact{}, id).Error
}
