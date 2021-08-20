package internals

import (
	tx "github.com/lamaorg/lama/Transaction"
	"github.com/lamaorg/lama/internals/primitives"
	"github.com/lamaorg/lama/storage"
	"math/big"
	"sync"
)

type iLamaBlockchain interface {
	GetChainID() *big.Int
	SetChainID()
	GetBlocks() []*primitives.Block
	localStorage() storage.LocalStorage
	GetProcessor() *Processor
	GetChainTip() (*primitives.BlockID, int64, error)
	GetBlockIDForHeight(height int64) (*primitives.BlockID, error)
	GetBranchType(id primitives.BlockID)
}

type Processor struct {
	GenesisID               *big.Int
	BlockRouterID           *big.Int
	OracleID                *big.Int
	BlockStorage            storage.LocalStorage
	txQueue                 []interface{}
	LamaChain               LamaBlockchain
	txChan                  chan txToProcess
	blockChan               chan blockToProcess
	registerNewTxChan       chan chan<- NewTx
	unregisterNewTxChan     chan chan<- NewTx
	registerTipChangeChan   chan chan<- TipChange
	unregisterTipChangeChan chan chan<- TipChange
	newTxChannels           map[chan<- NewTx]struct{}
	tipChangeChannels       map[chan<- TipChange]struct{}
	shutdownChan            chan struct{}

	wg sync.WaitGroup
}

type TipChange struct {
	BlockID *big.Int          // block ID of the main chain tip block
	Block   *primitives.Block // full block
	Source  string            // who sent the block that caused this change
	Connect bool              // true if the tip has been connected. false for disconnected
	More    bool              // true if the tip has been connected and more connections are expected
}

type blockToProcess struct {
	id         *big.Int          // block ID
	block      *primitives.Block // block to process
	source     string            // who sent it
	resultChan chan<- error      // channel to receive the result
}

type txToProcess struct {
	id         *big.Int        // transaction ID
	tx         *tx.Transaction // transaction to process
	source     string          // who sent it
	resultChan chan<- error    // channel to receive the result
}

type NewTx struct {
	TransactionID *big.Int        // transaction ID
	Transaction   *tx.Transaction // new transaction
	Source        string          // who sent it
}

func NewProcessor(genesisID *big.Int, blockStore storage.LocalStorage, txQueue []tx.Transaction, chain LamaBlockchain) *Processor {
	return &Processor{}
}
