package service

import (
	"database/sql"
	"pq/internal/app/model"
	"pq/internal/app/repository"
)

type Service struct {
	repo *repository.UserRepository
}

func NewService(db *sql.DB) *Service {
	return &Service{
		repo: repository.NewUserRepository(db),
	}
}

func (s *Service) CreateUser(user model.User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) GetUser(id int) (model.User, error) {
	return s.repo.GetUser(id)
}

func (s *Service) GetUsers() ([]model.User, error) {
	return s.repo.GetUsers()
}

func (s *Service) UpdateUser(id int, user model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *Service) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
