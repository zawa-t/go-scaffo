package onion

import (
	"embed"
	"path/filepath"

	"github.com/zawa-t/go-scaffo/src/template/config"
)

var (
	//go:embed template/*
	templates embed.FS

	templatePath = "template"
)

type Onion struct{}

func New() *Onion {
	return &Onion{}
}

func (o *Onion) LoadContents(basePath, appName string) []config.Content {
	return []config.Content{
		{Dir: filepath.Join(basePath, "cmd", "server"), Files: map[string]string{"main.go": "main.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "presentation"), Files: map[string]string{"rooter.go": "rooter.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "presentation", "handler"), Files: map[string]string{"user_handler.go": "handler.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "usecase"), Files: map[string]string{"user_usecase.go": "usecase.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "usecase", "service"), Files: map[string]string{"user_create_service.go": "user_create_service.go.tpl", "user_find_service.go": "user_find_service.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "domain", "user"), Files: map[string]string{"user.go": "domain.go.tpl", "user_repository.go": "domain_repository.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "infrastructure", "inmemory"), Files: map[string]string{"user_repository.go": "infrastructure.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "infrastructure", "db"), Files: map[string]string{"db.go": "db.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "dependency"), Files: map[string]string{"dependency.go": "dependency.go.tpl"}},
		{Dir: filepath.Join(basePath, "errors"), Files: map[string]string{"errors.go": "errors.go.tpl"}},
		{Dir: filepath.Join(basePath, "env"), Files: map[string]string{"env.go": "env.go.tpl"}},
		{Dir: filepath.Join(basePath, "env", "lang"), Files: map[string]string{"language.go": "language.go.tpl"}},
		{Dir: filepath.Join(basePath, "pkg", "logger"), Files: map[string]string{"logger.go": "logger.go.tpl"}},
	}
}

func (o *Onion) LoadTemplateConfig() config.Template {
	return config.Template{
		EmbededFiles: templates,
		Path:         templatePath,
	}
}
