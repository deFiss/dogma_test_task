package service

import (
	"dogma_test_task/internal"
	"dogma_test_task/internal/repository"
)

type User interface {
	GetUserList() ([]internal.User, error)
	AddUser(input internal.User) (int, error)
	GetUserById(id int) (internal.User, error)
	UpdateUser(id int, input internal.User) (internal.User, error)
	DeleteUser(id int) error
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUsersService(repos.Users),
	}
}
