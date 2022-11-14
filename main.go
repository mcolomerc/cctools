package main

import (
	"mcolomerc/cc-tools/cmd/cctools"
	"mcolomerc/cc-tools/pkg/log"
	"os"
)

func main() {
	lvl, ok := os.LookupEnv("LOG")
	if !ok {
		log.SetLevel(2) // default to the info level
	}
	//LOG=DEBUG
	if lvl == "DEBUG" {
		log.SetLevel(1)
	}
	cctools.Execute()

}
