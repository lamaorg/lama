package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lamaorg/lama/common"
	_ "github.com/lamaorg/lama/common"
)

func main() {
	addr := new(common.Address)
	a := addr.New()
	spew.Dump(a)
}
