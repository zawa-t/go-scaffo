package util

import (
	"fmt"
	"strings"
)

func ToUpperCase(input string) string {
	return strings.ToUpper(input)
}

func PrintBanner(appName string) {
	fmt.Printf("Welcome to %s CLI!\n", appName)
}
