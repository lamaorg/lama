package primitives

import (
	"github.com/lamaorg/lama/internals/dnaProof"
	"math/big"
)

type Coin struct {
	ID            *big.Int
	Proof         *dnaProof.Proof
	IsMinted      bool
	IsBurned      bool
	IsSpent       bool
	MintAtBlockID *big.Int
	ValidatedBy   string
	Signature     []byte
}

type Currency struct {
	Name        string
	Symbol      string
	Description string
	URL         string
	Issuer      string
	Metadata    map[string]string
	IssuedOn    int64
}
