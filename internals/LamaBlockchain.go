package internals

import (
	"bytes"
	"encoding/json"
	tx "github.com/lamaorg/lama/Transaction"
	"github.com/lamaorg/lama/common"
	"github.com/lamaorg/lama/internals/dnaProof"
	"github.com/lamaorg/lama/internals/primitives"
	"github.com/lunixbochs/struc"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

type LamaBlockchain struct {
	Genesis *Genesis
	Blocks  []*primitives.Block

	iLamaBlockchain
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

}

func init() {
	NewBlockchain()
}
