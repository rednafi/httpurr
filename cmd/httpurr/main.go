package main

import (
	"github.com/rednafi/httpurr/src"
)

// Ldflags filled by goreleaser
var Version string

func main() {
	src.Cli(&Version)

}
