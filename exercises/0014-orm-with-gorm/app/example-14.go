package app

import (
	"0014-orm-with-gorm/app/config"
	"0014-orm-with-gorm/app/model"
	"0014-orm-with-gorm/app/service"
)

// Example14 - struct to hold our configuration for the example
type Example14 struct {
	database *config.Database
	service  *service.Agenda
}

// NewExample14 - A constructor function so we can properly instantiate our
// sample code
func NewExample14() *Example14 {
	var app Example14
	app.database = config.NewDatabase()
	app.service = service.NewAgenda(app.database)
	return &app
}

func (example Example14) Insert(name string, address string) (uint, error) {
	return example.service.Insert(&model.Contact{
		Name: name,
		Addresses: []model.Address{
			{
				Street: address,
			},
		},
	})
}

func (example Example14) Update(contact *model.Contact) error {
	return example.service.Update(contact)
}

func (example Example14) Delete(id uint) error {
	return example.service.Delete(id)
}

func (example Example14) Find(id uint) (*model.Contact, error) {
	return example.service.Find(id)
}

func (example Example14) ListByName(q string) (*[]model.Contact, error) {
	return example.service.ListByName(q)
}

func (example Example14) ListByAddress(q string) (*[]model.Contact, error) {
	return example.service.ListByAddress(q)
}

func (example Example14) CloseDb() error {
	return example.database.Close()
}
