package project

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/zawa-t/go-scaffo/src/project/config"
	"github.com/zawa-t/go-scaffo/src/project/config/onion"
)

var (
	//go:embed config/onion/template/*
	templates embed.FS
)

type root struct {
	dir    string
	marker string
}

type Project struct {
	appName string
	root    root
}

func New(appName, pjtRootMarker string) (*Project, error) {
	rootDir, err := findProjectRootDir(pjtRootMarker)
	if err != nil {
		return nil, err
	}
	return &Project{
		appName: appName,
		root: root{
			dir:    rootDir,
			marker: pjtRootMarker,
		},
	}, nil
}

func (p *Project) AddConfiguration(archName string) error {
	basePath, err := p.basePath()
	if err != nil {
		return err
	}

	var configuration *config.Configuration
	switch archName {
	default:
		configuration = onion.LoadConfiguration(basePath, p.appName)
	}

	for _, content := range configuration.Contents {
		if err := os.MkdirAll(content.Dir, os.ModePerm); err != nil {
			return err
		}
		if err := p.makeFileAll(configuration.TemplatePath, content); err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) basePath() (basePath string, err error) {
	if p.appName != "" {
		basePath = filepath.Join(p.root.dir, p.appName)
		err := os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			return "", err
		}
	} else {
		basePath = p.root.dir
	}
	return
}

func (p *Project) getModuleName() (string, error) {
	if p.root.marker != "go.mod" {
		return "", errors.New("プロジェクト指標がgo.modではありません")
	}
	goModPath := filepath.Join(p.root.dir, p.root.marker)
	file, err := os.Open(goModPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", errors.New("module name not found in go.mod")
}

func (p *Project) makeFileAll(tmplPath string, content config.Structure) error {
	for filePath, tmplFile := range content.Files {
		file, err := os.Create(filepath.Join(content.Dir, filePath))
		if err != nil {
			return err
		}
		defer file.Close()

		tmpl, err := template.ParseFS(templates, filepath.Join(tmplPath, tmplFile))
		if err != nil {
			return err
		}

		moduleName, err := p.getModuleName()
		if err != nil {
			return err
		}

		var baseImportPath string
		if p.appName != "" {
			baseImportPath = fmt.Sprintf("%s/%s", moduleName, p.appName)
		} else {
			baseImportPath = moduleName
		}

		data := struct {
			AppName        string
			BaseImportPath string
		}{
			AppName:        p.appName,
			BaseImportPath: baseImportPath,
		}

		if err := tmpl.Execute(file, data); err != nil {
			return err
		}
	}
	return nil
}

func findProjectRootDir(pjtRootMarker string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// MEMO: pjtRootIndicatorが見つかるまで上位階層を順に探索
	for {
		path := filepath.Join(currentDir, pjtRootMarker)
		if _, err := os.Stat(path); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break // MEMO: ルートディレクトリに到達した場合はループ終了
		}
		currentDir = parentDir
	}
	return "", fmt.Errorf("%s not found", pjtRootMarker)
}
