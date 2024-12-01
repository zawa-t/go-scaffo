package usecase

import (
	"{{ .BaseImportPath }}/internal/domain/user"
	"{{ .BaseImportPath }}/internal/usecase/service"
)

type UserUsecase struct {
	service.UserCreate
	service.UserFind
}

func NewUserUsecase(sr user.UserRepository) *UserUsecase {
	return &UserUsecase{service.NewUserCreateService(sr), service.NewUserFindService(sr)}
}
