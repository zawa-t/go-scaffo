package env

import (
	"os"
)

var Lang = os.Getenv("LANG")

func getEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		slog.Warn("failed to get environment variable.", "name", name)
	}
	return v
}