package service

import (
	"{{ .BaseImportPath }}/internal/domain/user"
)

type UserFind interface {
	Find(id int) (*user.User, error)
}

type UserFindService struct {
	userRepository user.UserRepository
}

func NewUserFindService(sr user.UserRepository) *UserFindService {
	return &UserFindService{sr}
}

func (s *UserFindService) Find(id int) (*user.User, error) {
	return s.userRepository.FindBy(id)
}
