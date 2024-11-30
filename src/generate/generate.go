package generate

import (
	"github.com/zawa-t/go-scaffo/src/project"
)

func Scaffold(appName string) error {
	goModFile := "go.mod"
	pjt, err := project.New(appName, goModFile) // MEMO: go.mod ファイルがある場所をそのプロジェクトの root ディレクトリと定義
	if err != nil {
		return err
	}

	if err := pjt.AddConfiguration("onion"); err != nil {
		return err
	}
	return nil
}
