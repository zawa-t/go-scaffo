package main

import (
	"fmt"
	"os"

	"{{ .BaseImportPath }}/internal/command"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: {{ .AppName }} <command> [options]")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "help":
		command.Help()
	case "version":
		command.Version()
	case "{{ .CommandName }}":
		command.{{ .CommandName | capitalizeFirst }}(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		command.Help()
		os.Exit(1)
	}
}
