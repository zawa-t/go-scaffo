package template

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/zawa-t/go-scaffo/src/template/config"
)

type Data struct {
	AppName        string
	BaseImportPath string
	CommandName    string
}

func NewData(moduleName, appName string, commandName *string) *Data {
	var baseImportPath string
	if appName != "" {
		baseImportPath = fmt.Sprintf("%s/%s", moduleName, appName)
	} else {
		baseImportPath = moduleName
	}

	data := &Data{
		AppName:        appName,
		BaseImportPath: baseImportPath,
	}

	if commandName != nil {
		data.CommandName = *commandName
	}
	return data
}

type Template struct {
	template *template.Template
}

func New(tmplFileName string, tmplConfig config.Template) (*Template, error) {
	funcMap := template.FuncMap{
		"capitalizeFirst": CapitalizeFirst,
	}

	tmpl, err := template.New(tmplFileName).Funcs(funcMap).ParseFS(tmplConfig.EmbededFiles, filepath.Join(tmplConfig.Path, tmplFileName))
	if err != nil {
		return nil, err
	}
	return &Template{tmpl}, nil
}

func (t *Template) Execute(file io.Writer, data Data) error {
	if err := t.template.Execute(file, data); err != nil {
		return err
	}
	return nil
}

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
