package primitives

import (
	"bytes"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	tx "github.com/lamaorg/lama/Transaction"
	"github.com/lamaorg/lama/common"
	"github.com/lamaorg/lama/internals/dnaProof"
	"github.com/lamaorg/lama/storage"
	"github.com/lunixbochs/struc"
	"io/ioutil"
	"math/big"
	"os"
	"sync"
	"time"
)

type iLamaBlockchain interface {
	getChainID() *big.Int
	setChainID()
	getBlocks() []*iBlock
	localStorage() storage.LocalStorage
	getProcessor() *Processor
}

type LamaBlockchain struct {
	Genesis *Genesis
	Blocks  []*Block

	iLamaBlockchain
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
	BlockID *big.Int // block ID of the main chain tip block
	Block   *Block   // full block
	Source  string   // who sent the block that caused this change
	Connect bool     // true if the tip has been connected. false for disconnected
	More    bool     // true if the tip has been connected and more connections are expected
}

type blockToProcess struct {
	id         *big.Int     // block ID
	block      *Block       // block to process
	source     string       // who sent it
	resultChan chan<- error // channel to receive the result
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

func (L *LamaBlockchain) LoadGenesis() {

	data, err := ioutil.ReadFile("./genesis.json")

	if os.IsNotExist(err) {
		CreateGenesis()
		data, _ = ioutil.ReadFile("./genesis.json")
	}

	var G Genesis
	err = json.Unmarshal(data, &G)
	if err != nil {
		panic(err)
	}

	L.Genesis = &G

}

func (L *LamaBlockchain) GetChainID() *big.Int {
	return L.Genesis.ChainID
}

func CreateGenesis() {

	seed, _ := ioutil.ReadFile("./.seeds/seed32-10.txt")
	chainID := new(big.Int).SetBytes(seed)
	versionID := new(big.Int).SetBytes([]byte("v0.0.1"))

	g := new(Genesis)
	g.ChainID = chainID
	g.VersionID = versionID
	cb := tx.CreateNewInternalCoinbase()

	g.CoinbaseAddress = cb.Address
	g.Coin = "LLama"

	g.GenesisTime = time.Now().UnixNano()
	g.createProof()

	jsonized, _ := json.Marshal(g)
	err := ioutil.WriteFile("genesis.json", jsonized, 0644)
	if err != nil {
		panic(err)
	}

}

func (g *Genesis) createProof() {
	dur, _ := time.ParseDuration("5s")

	p := dnaProof.TemperProofParams{
		MutationRate:   0.005,
		ParseDuration:  dur,
		PopulationSize: 1000,
		MaxFitness:     0.5,
	}

	proof, hmac := dnaProof.GetProof(p, time.Now())
	var buf bytes.Buffer
	struc.Pack(&buf, hmac)
	g.BlockRoot = buf.String()
	var pBuf bytes.Buffer
	struc.Pack(&pBuf, proof)
	g.Proof = pBuf.String()
}

type Genesis struct {
	ChainID            *big.Int
	VersionID          *big.Int
	OriginatingHost    string
	OriginalValidators map[string]common.Address
	CoinbaseAddress    string
	Coin               string
	StakingCoin        string
	GenesisValidators  map[string]string
	BlockRoot          string
	GenesisTime        int64
	NextBlockHash      string
	Signature          string
	Proof              string
	commands           map[int]string
}

func NewBlockchain() {
	chain := new(LamaBlockchain)
	chain.LoadGenesis()
	spew.Dump(chain)
}

func init() {
	NewBlockchain()
}

func NewProcessor(genesisID *big.Int, blockStore storage.LocalStorage, txQueue []tx.Transaction, chain LamaBlockchain) *Processor {
	return &Processor{}
}
