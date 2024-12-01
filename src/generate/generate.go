package generate

import (
	"github.com/zawa-t/go-scaffo/src/project"
)

type Arg struct {
	AppName  string
	ArchName string
}

func Scaffold(arg Arg) error {
	pjt, err := project.New(arg.AppName, arg.ArchName)
	if err != nil {
		return err
	}

	if err := pjt.AddConfiguration(); err != nil {
		return err
	}
	return nil
}
