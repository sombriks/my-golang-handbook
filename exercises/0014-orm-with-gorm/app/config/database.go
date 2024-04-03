package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Database struct {
	Db *gorm.DB
}

func NewDatabase() *Database {
	var database Database
	var err error
	database.Db, err = gorm.Open(sqlite.Open("agenda.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = database.initDB()
	if err != nil {
		log.Fatal(err)
	}
	return &database
}

func (database *Database) Close() error {
	db, err := database.Db.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) initDB() error {
	return database.Db.Exec(`
		create table if not exists contacts(
		    id integer primary key,
		    name text not null,
		    created_at timestamp,
		    updated_at timestamp,
		    deleted_at timestamp               
		);
		create table if not exists addresses(
		    id integer primary key,
		    street text not null,
		    contact_id integer not null references contacts(id) on delete cascade,
		    created_at timestamp,
		    updated_at timestamp,
		    deleted_at timestamp
		);
	`).Error
}
