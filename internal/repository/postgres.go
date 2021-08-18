package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewPostgresql(cfg Config) (*sqlx.DB, error) {

	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password,
		),
	)
	if err != nil {
		return nil, err
	}

	db.Query("")

	return db, nil
}
