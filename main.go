package main

import (
	"os"

	"github.com/sophiabrandt/go-generate-social-banners/generate"
)

func main() {
	os.Exit(generate.CLI(os.Args[1:]))
}
