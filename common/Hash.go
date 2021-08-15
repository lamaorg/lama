package common

import (
	"golang.org/x/crypto/sha3"
	"math/big"
)

func HashBlock(b *big.Int) []byte {

	hash := sha3.New512()
	hash.Write(b.Bytes())
	return hash.Sum(nil)

}
