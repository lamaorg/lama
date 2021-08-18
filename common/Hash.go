package common

import (
	"golang.org/x/crypto/sha3"
	"hash"
	"math/big"
)

func (h Hashing) HashBlock(b *big.Int) []byte {

	h.hasher = sha3.New512()
	h.hasher.Reset()
	h.hasher.Write(b.Bytes())
	return h.hasher.Sum(nil)

}

type Hashing struct {
	hasher hash.Hash
	src    *big.Int
	dst    []byte
}
