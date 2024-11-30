package onion

import (
	"path/filepath"

	"github.com/zawa-t/go-scaffo/src/project/config"
)

var templatePath = "config/onion/template"

func LoadConfiguration(basePath string, appName string) *config.Configuration {
	return &config.Configuration{
		TemplatePath: templatePath,
		Contents: config.Contents{
			{Dir: filepath.Join(basePath, "cmd", "server"), Files: map[string]string{"main.go": "main.go.tpl"}},
			{Dir: filepath.Join(basePath, "internal", "presentation"), Files: map[string]string{"rooter.go": "rooter.go.tpl"}},
			{Dir: filepath.Join(basePath, "internal", "presentation", "handler"), Files: map[string]string{"user_handler.go": "handler.go.tpl"}},
			{Dir: filepath.Join(basePath, "internal", "usecase"), Files: map[string]string{"user_usecase.go": "usecase.go.tpl"}},
			{Dir: filepath.Join(basePath, "internal", "usecase", "service"), Files: map[string]string{"user_create_service.go": "user_create_service.go.tpl", "user_find_service.go": "user_find_service.go.tpl"}},
			{Dir: filepath.Join(basePath, "internal", "domain", "user"), Files: map[string]string{"user.go": "domain.go.tpl", "user_repository.go": "domain_repository.go.tpl"}},
			{Dir: filepath.Join(basePath, "internal", "infrastructure", "inmemory"), Files: map[string]string{"user_repository.go": "infrastructure.go.tpl"}},
			{Dir: filepath.Join(basePath, "internal", "infrastructure", "db"), Files: map[string]string{"db.go": "db.go.tpl"}},
			{Dir: filepath.Join(basePath, "internal", "dependency"), Files: map[string]string{"dependency.go": "dependency.go.tpl"}},
			{Dir: filepath.Join(basePath, "pkg", "logger"), Files: map[string]string{"logger.go": "logger.go.tpl"}},
		},
	}
}
