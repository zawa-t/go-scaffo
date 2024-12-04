package command

import (
	"flag"
	"fmt"
)

func {{ .CommandName | capitalizeFirst }} (args []string) {
	fs := flag.NewFlagSet("{{ .CommandName }}", flag.ExitOnError)

	var option string
	fs.StringVar(&option, "option", "", "An example option")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("Failed to parse flags: %v\n", err)
		fs.Usage()
		return
	}

	fmt.Printf("Executing {{ .CommandName }} with option: %s\n", option)
}
