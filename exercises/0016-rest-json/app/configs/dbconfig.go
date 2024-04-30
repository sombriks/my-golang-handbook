package configs

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	db *goqu.Database
}

func NewDBConfig() (*DBConfig, error) {
	var dbConfig DBConfig
	con, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost/todo")
	if err != nil {
		return nil, err
	}
	dbConfig.db = goqu.New("postgres", con)
	return &dbConfig, nil
}
