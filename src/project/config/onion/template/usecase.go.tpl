package usecase

import (
	"{{ .ModuleName }}/{{ .AppName }}/internal/usecase/service"
)

type UserUsecase struct {
	service.UserCreate
	service.UserFind
}

func NewUserUsecase(uc service.UserCreate, uf service.UserFind) UserUsecase {
	return UserUsecase{uc, uf}
}
