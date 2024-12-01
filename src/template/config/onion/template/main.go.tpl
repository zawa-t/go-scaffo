package main

import (
	"log"
	"net/http"

	"{{ .BaseImportPath }}/internal/dependency"
	"{{ .BaseImportPath }}/internal/presentation"
)

func main() {
	di := dependency.NewInjection().NewAppHandler()

	presentation.NewRouter(di)

	const port = ":8100"
	log.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
