package presentation

import (
	"net/http"

	"{{ .BaseImportPath }}/internal/injection"
)

func NewRouter(di *injection.Dependency) {
	http.HandleFunc("/user/create", di.UserHandler.CreateUser)
	http.HandleFunc("/user", di.UserHandler.GetUser)
}
