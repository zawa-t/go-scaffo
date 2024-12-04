package command

import "fmt"

func Help() {
	fmt.Println("Available commands:")
	fmt.Println("  help     Show this help message")
	fmt.Println("  version  Show application version")
	fmt.Println("  {{ .CommandName }}    Execute the {{ .CommandName }} command")
}

func Version() {
	fmt.Println("{{ .AppName }} version 1.0.0")
}