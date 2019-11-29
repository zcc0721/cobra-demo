package main

import (
	"runtime"

	cmd "github.com/zcc0721/cobra-demo/cmd/democli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
