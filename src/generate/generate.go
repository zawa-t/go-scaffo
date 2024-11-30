package generate

import (
	"os"

	"github.com/zawa-t/go-scaffo/src/project"
	"github.com/zawa-t/go-scaffo/src/project/config/onion"
)

func Scaffold(appName string) error {
	goModFile := "go.mod"
	pjt, err := project.New(appName, goModFile)
	if err != nil {
		return err
	}

	basePath, err := pjt.BasePath()
	if err != nil {
		return err
	}

	tmplPath := "config/onion/template"

	configurations := onion.LoadConfigurations("onion", basePath, pjt.AppName)
	for _, configuration := range configurations {
		if err := os.MkdirAll(configuration.Dir, os.ModePerm); err != nil {
			return err
		}

		if err := pjt.MakeFileAll(tmplPath, configuration.Dir, configuration.Files); err != nil {
			return err
		}
	}
	return nil
}
