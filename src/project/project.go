package project

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	AppName string
	root    root
}

func New(appName, pjtRootMarker string) (*Project, error) {
	rootDir, err := findProjectRootDir(pjtRootMarker) // MEMO: go.mod ファイルがある場所をそのプロジェクトの root ディレクトリと定義
	if err != nil {
		return nil, err
	}
	return &Project{
		AppName: appName,
		root: root{
			dir:    rootDir,
			marker: pjtRootMarker,
		},
	}, nil
}

func getModuleName(gomod io.Reader) (string, error) {
	scanner := bufio.NewScanner(gomod)
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

func (p *Project) BasePath() (basePath string, err error) {
	if p.AppName != "" {
		basePath = filepath.Join(p.root.dir, p.AppName)
		err := os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			return "", err
		}
	} else {
		basePath = p.root.dir
	}
	return
}

func (p *Project) moduleName() (string, error) {
	if p.root.marker != "go.mod" {
		return "", errors.New("プロジェクト指標がgo.modではありません")
	}
	goModPath := filepath.Join(p.root.dir, p.root.marker)
	file, err := os.Open(goModPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	moduleName, err := getModuleName(file)
	if err != nil {
		return "", err
	}
	return moduleName, nil
}

func (p *Project) MakeFileAll(tmplPath, dir string, files map[string]string) error {
	for filePath, tmplFile := range files {
		file, err := os.Create(filepath.Join(dir, filePath))
		if err != nil {
			return err
		}
		defer file.Close()

		tmpl, err := template.ParseFS(templates, filepath.Join(tmplPath, tmplFile))
		if err != nil {
			return err
		}

		moduleName, err := p.moduleName()
		if err != nil {
			return err
		}

		type Data struct {
			AppName    string
			ModuleName string
		}

		data := Data{
			AppName:    p.AppName,
			ModuleName: moduleName,
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
