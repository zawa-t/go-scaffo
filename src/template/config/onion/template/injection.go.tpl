package injection

import (
	"{{ .BaseImportPath }}/internal/infrastructure/inmemory"
	"{{ .BaseImportPath }}/internal/presentation/handler"
	"{{ .BaseImportPath }}/internal/usecase"
)

type Dependency struct {
	*appHandler
}

func NewDependency() *Dependency {
	d := new(Dependency)
	d.appHandler = newAppHandler(d)
	return d
}

type appHandler struct {
	*handler.UserHandler
}

func newAppHandler(d *Dependency) *appHandler {
	return &appHandler{
		UserHandler: d.newUserHandler(),
	}
}

func (d *Dependency) newUserHandler() *handler.UserHandler {
	return handler.NewUserHandler(d.newUserUsecase())
}

func (d *Dependency) newUserUsecase() *usecase.UserUsecase {
	return usecase.NewUserUsecase(d.newInMemoryUserRepository())
}

func (d *Dependency) newInMemoryUserRepository() *inmemory.UserRepository {
	return inmemory.NewUserRepository()
}
