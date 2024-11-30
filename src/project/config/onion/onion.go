package onion

import (
	"path/filepath"

	"github.com/zawa-t/go-scaffo/src/project"
)

var TemplatePath = "config/onion/template"

func LoadConfigurations(configName string, basePath string, appName string) project.Configurations {
	return project.Configurations{
		{Dir: filepath.Join(basePath, "cmd", appName), Files: map[string]string{"main.go": "main.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "handler"), Files: map[string]string{"user_handler.go": "handler.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "usecase"), Files: map[string]string{"user_usecase.go": "usecase.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "usecase", "service"), Files: map[string]string{"user_create_service.go": "user_create_service.go.tpl", "user_find_service.go": "user_find_service.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "domain", "user"), Files: map[string]string{"user.go": "domain.go.tpl", "user_repository.go": "domain_repository.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "infrastructure"), Files: map[string]string{"user_repository.go": "infrastructure.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "infrastructure", "db"), Files: map[string]string{"db.go": "db.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "dependency"), Files: map[string]string{"dependency.go": "dependency.go.tpl"}},
		{Dir: filepath.Join(basePath, "pkg", "logger"), Files: map[string]string{"logger.go": "logger.go.tpl"}},
	}
}
