package repository

import (
	"dogma_test_task/internal"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	GetUserList() ([]internal.User, error)
	AddUser(input internal.User) (int, error)
	GetUserById(id int) (internal.User, error)
	UpdateUser(id int, input internal.User) (internal.User, error)
	DeleteUser(id int) error
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUserPostgres(db),
	}
}
