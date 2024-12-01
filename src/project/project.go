package project

import (
	"bufio"
	"embed"
	"errors"
	"fmt"

	htmlTemplate "html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/zawa-t/go-scaffo/src/template"
	"github.com/zawa-t/go-scaffo/src/template/config/cli"
	"github.com/zawa-t/go-scaffo/src/template/config/onion"
)

type root struct {
	dir        string
	moduleName string
}

type Project struct {
	root    root
	appName string
	loader  func(basePath string, appName string) *template.Config
}

func New(appName, archName string) (*Project, error) {
	rootMarker := "go.mod"
	rootDir, err := findProjectRootDir(rootMarker) // MEMO: go.mod ファイルがある場所をそのプロジェクトの root ディレクトリと定義
	if err != nil {
		return nil, err
	}
	moduleName, err := getModuleName(rootDir, rootMarker)
	if err != nil {
		return nil, err
	}
	pjt := &Project{
		root: root{
			dir:        rootDir,
			moduleName: moduleName,
		},
		appName: appName,
	}

	switch archName {
	case "cli":
		pjt.loader = cli.LoadConfiguration
	default:
		pjt.loader = onion.LoadConfiguration
	}

	return pjt, nil
}

func (p *Project) AddConfiguration() error {
	basePath, err := p.basePath()
	if err != nil {
		return err
	}

	config := p.loader(basePath, p.appName)
	for _, content := range config.Contents {
		if err := os.MkdirAll(content.Dir, os.ModePerm); err != nil {
			return err
		}
		if err := p.makeFileAll(config.EmbededFiles, content); err != nil {
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

func (p *Project) makeFileAll(embededTemplateFiles embed.FS, content template.Content) error {
	for filePath, tmplFilePath := range content.Files {
		file, err := os.Create(filepath.Join(content.Dir, filePath))
		if err != nil {
			return err
		}
		defer file.Close()

		tmpl, err := htmlTemplate.ParseFS(embededTemplateFiles, tmplFilePath)
		if err != nil {
			return err
		}

		var baseImportPath string
		if p.appName != "" {
			baseImportPath = fmt.Sprintf("%s/%s", p.root.moduleName, p.appName)
		} else {
			baseImportPath = p.root.moduleName
		}

		data := template.Data{
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

func getModuleName(rootDir, rootMarker string) (string, error) {
	goModPath := filepath.Join(rootDir, rootMarker)
	file, err := os.Open(goModPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return "", err
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", errors.New("module name not found in go.mod")
}
