package configs

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
	"log"
)

type DBConfig struct {
	DB *goqu.Database
}

// NewDBConfig - database configurations (connection, initial scripts, etc
func NewDBConfig() (*DBConfig, error) {
	var dbConfig DBConfig
	con, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost/todo?sslmode=disable")
	if err != nil {
		return nil, err
	}
	dbConfig.DB = goqu.New("postgres", con)
	dbConfig.DB.Logger(log.Default())
	return &dbConfig, nil
}
