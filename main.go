package main

import (
	"github.com/lamaorg/lama/api"
	"github.com/lamaorg/lama/cmd"
)

func main() {

	api.Bootstrapper()

	cmd.Execute()
}
