package primitives

import (
	"github.com/lamaorg/lama/Transaction"
	"math/big"
)

const HEXTYPE = "INTERNAL"

type Block struct {
	ChainID        *big.Int
	Time           int64
	PrevRootHash   []byte
	Hash           []byte
	ValidatorAddr  string
	TransactionIDs map[string]Transaction.Transaction
}

type Header struct {
	BlockIndex int
}
