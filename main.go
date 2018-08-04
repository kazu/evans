package main

import (
	"os"

	"github.com/kazu/evans/adapter/controller"
	"github.com/kazu/evans/meta"
)

func main() {
	os.Exit(controller.NewCLI(
		meta.AppName,
		meta.Version.String(),
		controller.NewBasicUI(),
	).Run(os.Args[1:]))
}
