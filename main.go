package main

import (
	"github.com/mprimi/golang-cli-template/cmd"
	"os"
)

func main() {
	os.Exit(cmd.Run(os.Args[1:]))
}
