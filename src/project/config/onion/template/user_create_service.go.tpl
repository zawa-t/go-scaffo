package service

import (
	"{{ .ModuleName }}/{{ .AppName }}/internal/domain/user"
)

type UserCreate interface {
	Create(name string) error
}

type userCreateService struct {
	userRepository user.UserRepository
}

func NewUserCreateService(sr user.UserRepository) UserCreate {
	return &userCreateService{sr}
}

func (s *userCreateService) Create(name string) error {
	user := user.User{
		ID:   1,
		Name: name,
	}
	return s.userRepository.Save(user)
}
