package primitives

import "github.com/lamaorg/lama/storage"

type BlockchainIterator struct {
	CurrentBlock  *Block
	CurrentHeader *Header
	*storage.LocalStorage
}
