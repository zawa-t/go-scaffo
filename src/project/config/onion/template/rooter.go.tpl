package presentation

import (
	"net/http"

	"{{ .BaseImportPath }}/internal/dependency"
)

func NewRouter(di *dependency.AppHandler) {
	http.HandleFunc("/user/create", di.UserHandler.CreateUser)
	http.HandleFunc("/user", di.UserHandler.GetUser)
}
