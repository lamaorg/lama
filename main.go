package main

import (
	"github.com/lamaorg/lama/cmd"
	_ "github.com/lamaorg/lama/common"
	"github.com/lamaorg/lama/internals/primitives"
	"math/big"
)

var chainID *big.Int

func main() {

	if chainID == nil {
		chainID = primitives.GenerateChainID()
	}
	cmd.Execute()
}
