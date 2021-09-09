package main

import (
	"github.com/prompt-cash/js-transpiler/cmd"
	"runtime"
)

func main() {
	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd.Execute()
}
