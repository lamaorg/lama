package primitives

import "math/big"

const HEXTYPE = "INTERNAL"

var EmptyRootHash = []byte("fe127cd046bf452c49ad2e659dd49650b7e672a0f8348f819fab3baed47827b8")

type Block struct {
	ChainID       *big.Int
	Time          int64
	PrevRootHash  []byte
	Hash          []byte
	ValidatorAddr string
}

type Header struct{}
