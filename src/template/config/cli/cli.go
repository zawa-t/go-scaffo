package cli

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

type CLI struct{}

func New() *CLI {
	return &CLI{}
}

func (c *CLI) LoadContents(basePath, appName string) []config.Content {
	return []config.Content{
		{Dir: filepath.Join(basePath, "cmd", "cli"), Files: map[string]string{"main.go": "main.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "command"), Files: map[string]string{"root.go": "root.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "command"), Files: map[string]string{"command.go": "command.go.tpl"}},
		{Dir: filepath.Join(basePath, "internal", "config"), Files: map[string]string{"config.go": "config.go.tpl"}},
		{Dir: filepath.Join(basePath, "pkg", "logger"), Files: map[string]string{"logger.go": "logger.go.tpl"}},
		{Dir: filepath.Join(basePath, "pkg", "util"), Files: map[string]string{"util.go": "util.go.tpl"}},
		{Dir: filepath.Join(basePath), Files: map[string]string{"README.md": "README.md.tpl"}},
	}
}

func (c *CLI) LoadTemplateConfig() config.Template {
	return config.Template{
		EmbededFiles: templates,
		Path:         templatePath,
	}
}
