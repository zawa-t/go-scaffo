package generate

import (
	"embed"
	"html/template"
	"os"
	"path/filepath"
)

var (
	//go:embed template/*
	templates embed.FS
)

func Scaffold(appName string) error {
	basePath := "."
	if appName != "" {
		basePath = filepath.Join(basePath, appName)
		err := os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	dirs := map[string]map[string]string{
		filepath.Join(basePath, "cmd", appName):                     {"main.go": "main.go.tpl"},
		filepath.Join(basePath, "internal", "handler"):              {"sample_handler.go": "handler.go.tpl"},
		filepath.Join(basePath, "internal", "usecase"):              {"sample_usecase.go": "usecase.go.tpl"},
		filepath.Join(basePath, "internal", "domain"):               {"sample.go": "domain.go.tpl", "sample_repository.go": "domain_repository.go.tpl"},
		filepath.Join(basePath, "internal", "infrastructure"):       {"sample_repository.go": "infrastructure.go.tpl"},
		filepath.Join(basePath, "internal", "infrastructure", "db"): {"db.go": "db.go.tpl"},
		filepath.Join(basePath, "pkg", "logger"):                    {"logger.go": "logger.go.tpl"},
	}

	for dir, files := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}

		for filePath, tmplFile := range files {
			file, err := os.Create(filepath.Join(dir, filePath))
			if err != nil {
				return err
			}
			defer file.Close()

			// tmpl, err := template.ParseFiles(tmplPath)
			tmpl, err := template.ParseFS(templates, filepath.Join("template", tmplFile))
			if err != nil {
				return err
			}

			type Data struct {
				AppName string
			}

			data := Data{AppName: appName}

			if err := tmpl.Execute(file, data); err != nil {
				return err
			}
		}
	}
	return nil
}
