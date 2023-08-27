package main

import (
	"github.com/rednafi/httpurr/src"
)

// Ldflags filled by goreleaser
var version string

func main() {
	src.Cli(version)

}
