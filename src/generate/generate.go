package generate

import (
	"github.com/zawa-t/go-scaffo/src/project"
	"github.com/zawa-t/go-scaffo/src/template/config/cli"
	"github.com/zawa-t/go-scaffo/src/template/config/onion"
)

type Arg struct {
	AppName     string
	ArchName    string
	CommandName string
}

func Scaffold(arg Arg) error {
	var Loader project.Loader
	switch arg.ArchName {
	case "cli":
		Loader = cli.New()
	default:
		Loader = onion.New()
	}

	pjt, err := project.New(arg.AppName, arg.ArchName, arg.CommandName, Loader)
	if err != nil {
		return err
	}

	if err := pjt.AddConfiguration(); err != nil {
		return err
	}
	return nil
}
