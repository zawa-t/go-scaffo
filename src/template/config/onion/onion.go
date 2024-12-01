package onion

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
			{Dir: filepath.Join(basePath, "cmd", "server"), Files: map[string]string{"main.go": filepath.Join(templatePath, "main.go.tpl")}},
			{Dir: filepath.Join(basePath, "internal", "presentation"), Files: map[string]string{"rooter.go": filepath.Join(templatePath, "rooter.go.tpl")}},
			{Dir: filepath.Join(basePath, "internal", "presentation", "handler"), Files: map[string]string{"user_handler.go": filepath.Join(templatePath, "handler.go.tpl")}},
			{Dir: filepath.Join(basePath, "internal", "usecase"), Files: map[string]string{"user_usecase.go": filepath.Join(templatePath, "usecase.go.tpl")}},
			{Dir: filepath.Join(basePath, "internal", "usecase", "service"), Files: map[string]string{"user_create_service.go": filepath.Join(templatePath, "user_create_service.go.tpl"), "user_find_service.go": filepath.Join(templatePath, "user_find_service.go.tpl")}},
			{Dir: filepath.Join(basePath, "internal", "domain", "user"), Files: map[string]string{"user.go": filepath.Join(templatePath, "domain.go.tpl"), "user_repository.go": filepath.Join(templatePath, "domain_repository.go.tpl")}},
			{Dir: filepath.Join(basePath, "internal", "infrastructure", "inmemory"), Files: map[string]string{"user_repository.go": filepath.Join(templatePath, "infrastructure.go.tpl")}},
			{Dir: filepath.Join(basePath, "internal", "infrastructure", "db"), Files: map[string]string{"db.go": filepath.Join(templatePath, "db.go.tpl")}},
			{Dir: filepath.Join(basePath, "internal", "dependency"), Files: map[string]string{"dependency.go": filepath.Join(templatePath, "dependency.go.tpl")}},
			{Dir: filepath.Join(basePath, "pkg", "logger"), Files: map[string]string{"logger.go": filepath.Join(templatePath, "logger.go.tpl")}},
		},
	}
}
