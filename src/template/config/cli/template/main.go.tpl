package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "World", "Name to greet")
	flag.Parse()

	if len(os.Args) > 1 {
		fmt.Printf("Hello, %s!\n", os.Args[1])
	} else {
		fmt.Printf("Hello, %s!\n", *name)
	}
}
