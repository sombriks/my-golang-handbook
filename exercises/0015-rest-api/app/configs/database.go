package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db *gorm.DB
}

// NewDatabase - provision a ready to business database instance
func NewDatabase() (*Database, error) {

	var err error
	var database Database

	database.Db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	err = database.initDb()
	if err != nil {
		return nil, err
	}
	return &database, nil
}

func (db *Database) initDb() error {
	return db.Db.Exec(`
		create table if not exists todos (
		    id integer primary key,
		    description text not null,
		    done boolean not null default false,
		    created timestamp not null default CURRENT_TIMESTAMP
		);
	`).Error
}
