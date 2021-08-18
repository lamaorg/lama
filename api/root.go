package api

import (
	network "github.com/lamaorg/lama/Network"
	"github.com/lamaorg/lama/Transaction"
	"github.com/lamaorg/lama/common"
	"github.com/lamaorg/lama/internals/dnaProof"
	"github.com/lamaorg/lama/storage"
)

/*
	lama/api package
	to make it easier to use the
	functions in many contexts
*/

type API struct {
	Keys        common.SecureKeys
	Hash        common.Hashing
	HexTools    common.HexTools
	HexOperator common.HexOperator
	Addresses   common.RawAddress
	Proofing    dnaProof.Proof
	ChainConfig map[string]string
	Storage     storage.LocalStorage
	Transaction Transaction.Base
	Network     network.LLxPosNetwork
}

var GetApi *API

func init() {
	GetApi = new(API)
}
