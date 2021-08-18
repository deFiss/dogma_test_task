package service

import (
	"dogma_test_task/internal"
	"dogma_test_task/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) GetUserList() ([]internal.User, error) {
	return s.repo.GetUserList()
}

func (s *UsersService) AddUser(input internal.User) (int, error) {
	return s.repo.AddUser(input)
}

func (s *UsersService) GetUserById(id int) (internal.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UsersService) UpdateUser(id int, input internal.User) (internal.User, error) {
	return s.repo.UpdateUser(id, input)
}

func (s *UsersService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}