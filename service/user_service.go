package service

import (
	"todolist-graphql/entity"
	"todolist-graphql/repository"
)

type UserService interface {
	FindAll() ([]*entity.User, error)
	FindById(id int32) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) FindAll() ([]*entity.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) FindById(id int32) (*entity.User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Create(user *entity.User) (*entity.User, error) {
	user, err := s.repository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
