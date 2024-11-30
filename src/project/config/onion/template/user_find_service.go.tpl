package service

import (
	"{{ .ModuleName }}/{{ .AppName }}/internal/domain/user"
)

type UserFind interface {
	Find(id int) (*user.User, error)
}

type userFindService struct {
	userRepository user.UserRepository
}

func NewUserFindUsecase(sr user.UserRepository) UserFind {
	return &userFindService{sr}
}

func (s *userFindService) Find(id int) (*user.User, error) {
	return s.userRepository.FindBy(id)
}
