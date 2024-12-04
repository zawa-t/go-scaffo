package project

import (
	"bufio"
	"errors"
	"fmt"

	"os"
	"path/filepath"
	"strings"

	"github.com/zawa-t/go-scaffo/src/template"
	"github.com/zawa-t/go-scaffo/src/template/config"
	"github.com/zawa-t/go-scaffo/src/template/config/cli"
	"github.com/zawa-t/go-scaffo/src/template/config/onion"
)

type Project struct {
	root        string
	moduleName  string
	appName     string
	commandName *string
	arch        string
	loadConfig  func(basePath, appName string) *config.Configuration
}

func New(appName, archName, commandName string) (*Project, error) {
	if archName == "cli" && commandName == "" {
		return nil, errors.New("CLI構成の場合、コマンド名は必須です")
	}
	if archName != "cli" && commandName != "" {
		return nil, errors.New("CLI構成以外の場合、コマンド名は指定できません")
	}

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
		root:       rootDir,
		moduleName: moduleName,
		appName:    appName,
		arch:       archName,
	}

	switch archName {
	case "cli":
		pjt.commandName = &commandName
		pjt.loadConfig = cli.LoadConfiguration
	default:
		pjt.loadConfig = onion.LoadConfiguration
	}

	return pjt, nil
}

func (p *Project) AddConfiguration() error {
	basePath, err := p.basePath()
	if err != nil {
		return err
	}
	config := p.loadConfig(basePath, p.appName)
	for _, content := range config.Contents {
		if err := createDirectory(content.Dir); err != nil {
			return err
		}
		if err := p.makeFiles(content.Dir, content.Files, config.Template); err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) basePath() (basePath string, err error) {
	if p.appName != "" {
		basePath = filepath.Join(p.root, p.appName)
		err := os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			return "", err
		}
	} else {
		basePath = p.root
	}
	return
}

func (p *Project) makeFiles(dir string, files map[string]string, tmplConfig config.Template) error {
	for fileName, tmplFileName := range files {
		if err := func() error {
			file, err := os.Create(filepath.Join(dir, fileName))
			if err != nil {
				return err
			}
			defer file.Close()

			tmpl, err := template.New(tmplConfig, tmplFileName)
			if err != nil {
				return err
			}

			data := template.NewData(p.moduleName, p.appName, p.commandName)
			if err := tmpl.Execute(file, *data); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return err
		}
	}
	return nil
}

// createDirectory ensures the directory exists.
func createDirectory(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
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
