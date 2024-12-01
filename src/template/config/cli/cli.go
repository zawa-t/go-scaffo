package cli

import (
	"embed"
	"path/filepath"

	"github.com/zawa-t/go-scaffo/src/template"
)

var (
	//go:embed template/*
	templates embed.FS

	templatePath = "template"
)

func LoadConfiguration(basePath string, appName string) *template.Config {
	return &template.Config{
		EmbededFiles: templates,
		Contents: []template.Content{
			{Dir: filepath.Join(basePath, "cmd", "cli"), Files: map[string]string{"main.go": filepath.Join(templatePath, "main.go.tpl")}},
		},
	}
}
