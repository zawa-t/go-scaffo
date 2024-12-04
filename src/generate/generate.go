package generate

import (
	"github.com/zawa-t/go-scaffo/src/project"
)

type Arg struct {
	AppName     string
	ArchName    string
	CommandName string
}

func Scaffold(arg Arg) error {
	pjt, err := project.New(arg.AppName, arg.ArchName, arg.CommandName)
	if err != nil {
		return err
	}

	if err := pjt.AddConfiguration(); err != nil {
		return err
	}
	return nil
}
