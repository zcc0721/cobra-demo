package main

import (
	"runtime"

	cmd "github.com/cobra-demo/cmd/democli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
