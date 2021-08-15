package primitives

import (
	"crypto/rand"
	"github.com/lamaorg/lama/storage"
	"math"
	"math/big"
)

type iLamaBlockchain interface {
	getChainID() *big.Int
	setChainID()
	getBlocks() []*iBlock
	localStorage() storage.LocalStorage
}

type LamaBlockchain struct {
	ChainID *big.Int
	Blocks  []*Block
}

func GenerateChainID() *big.Int {

	b := new(big.Int).SetUint64(math.MaxUint64)
	b2, err := rand.Int(rand.Reader, b)
	if err != nil {
		return nil
	}
	return b2

}

func (L *LamaBlockchain) setChainID() {
	L.ChainID = GenerateChainID()
}
