package env

import (
	"os"
)

var Lang = os.Getenv("LANG")
