package main

import (
	"log"
	"net/http"

	"{{ .BaseImportPath }}/internal/injection"
	"{{ .BaseImportPath }}/internal/presentation"
)

func main() {
	di := injection.NewDependency()

	presentation.NewRouter(di)

	const port = ":8100"
	log.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
