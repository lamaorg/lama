package cmd

/*
	special cmd to generate a genesis file payload
*/

import (
	"bytes"
	"encoding/json"
	network "github.com/lamaorg/lama/Network"
	"github.com/lamaorg/lama/Transaction"
	"github.com/lamaorg/lama/common"
	"github.com/lamaorg/lama/internals/dnaProof"
	"github.com/lamaorg/lama/internals/primitives"
	"github.com/lunixbochs/struc"
	"github.com/spf13/cobra"
	"io/ioutil"
	"math/big"
	"time"
)

const EMPTYBLOCKROOT = "Llx000000002069732061207368697474792073656e74656e636520746f2074657374Lx41444452"

var (
	GenesisCMD = &cobra.Command{
		Use:   "genesis",
		Short: "Create and manage genesis block",
		Long:  "Create and manage genesis block (need generated operator keys)",
		Run: func(cmd *cobra.Command, args []string) {
			CreateGenesis()
		},
	}
)

func CreateGenesis() {

	seed, _ := ioutil.ReadFile("./.seeds/seed32-10.txt")
	chainID := new(big.Int).SetBytes(seed)
	versionID := new(big.Int).SetBytes([]byte("v0.0.1"))

	g := new(Genesis)
	g.ChainID = chainID
	g.VersionID = versionID
	cb := Transaction.CreateNewInternalCoinbase()

	g.CoinbaseAddress = cb.Address
	g.Coin = cb.Currency

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
	Coin               *primitives.Currency
	StakingCoin        *primitives.Currency
	GenesisValidators  []SuperValidators
	BlockRoot          string
	GenesisTime        int64
	NextBlockHash      string
	Signature          string
	Proof              string
	commands           map[int]string
}

type SuperValidators struct {
	Address   string
	account   network.Node
	childrens []network.Node
	Staked    int
}
