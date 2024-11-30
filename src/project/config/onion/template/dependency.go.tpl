package dependency

import (
	"{{ .BaseImportPath }}/internal/infrastructure"
	"{{ .BaseImportPath }}/internal/presentation/handler"
	"{{ .BaseImportPath }}/internal/usecase"
)

type Injection struct{}

func NewInjection() *Injection {
	return &Injection{}
}

type AppHandler struct {
	UserHandler *handler.UserHandler
}

func (i *Injection) NewAppHandler() *AppHandler {
	return &AppHandler{
		UserHandler: i.newUserHandler(),
	}
}

func (i *Injection) newUserHandler() *handler.UserHandler {
	return handler.NewUserHandler(i.newUserUsecase())
}

func (i *Injection) newUserUsecase() *usecase.UserUsecase {
	return usecase.NewUserUsecase(i.newInMemoryUserRepository())
}

func (i *Injection) newInMemoryUserRepository() *infrastructure.InMemoryUserRepository {
	return infrastructure.NewInMemoryUserRepository()
}
