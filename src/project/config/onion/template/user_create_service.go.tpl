package service

import (
	"{{ .BaseImportPath }}/internal/domain/user"
)

type UserCreate interface {
	Create(name string) error
}

type UserCreateService struct {
	userRepository user.UserRepository
}

func NewUserCreateService(sr user.UserRepository) *UserCreateService {
	return &UserCreateService{sr}
}

func (s *UserCreateService) Create(name string) error {
	user := user.User{
		ID:   1,
		Name: name,
	}
	return s.userRepository.Save(user)
}
