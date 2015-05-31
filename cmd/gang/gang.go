package main

import (
	_ "fmt"
	"github.com/ysugimoto/gang"
	"github.com/ysugimoto/go-cliargs"
)

func main() {
	args := cliarg.NewArguments()
	args.Alias("h", "help", false)
	args.Parse()

	runner := gang.NewRunner(args)
	runner.Run()
}
