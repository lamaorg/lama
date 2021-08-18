package network

import (
	"encoding/json"
	"errors"
	"github.com/lamaorg/lama/common"
	"github.com/lamaorg/lama/internals/dag"
	"github.com/lamaorg/lama/internals/primitives"

	"math/rand"

	"math/big"
	"time"
)

type LLxPosNetwork struct {
	Blockchain     *primitives.LamaBlockchain
	BlockchainHead *big.Int
	Validators     []*Node
	DAG            *dag.DAG
}

type Node struct {
	Stake   uint
	Account string
}

func (n LLxPosNetwork) GenerateNewBlock(Validator *Node) (*primitives.LamaBlockchain, *big.Int, error) {

	if err := n.ValidateChain(); err != nil {
		Validator.Stake -= 10
		return n.Blockchain, n.BlockchainHead, err
	}
	blockLen := len(n.Blockchain.Blocks)
	currentTime := time.Now().UnixNano()
	prevroot := n.Blockchain.Blocks[blockLen-1].Hash

	if blockLen == 0 {
		//need genesis
		prevroot = primitives.EmptyRootHash
	}

	newBlock := &primitives.Block{
		ChainID:       n.Blockchain.ChainID,
		Time:          currentTime,
		PrevRootHash:  prevroot,
		Hash:          []byte(""),
		ValidatorAddr: Validator.Account,
	}
	newBlockBytes, _ := json.Marshal(newBlock)
	blockBytes := new(big.Int).SetBytes(newBlockBytes).SetBytes(prevroot)
	var hasher common.Hashing
	hash := hasher.HashBlock(blockBytes)
	newBlock.Hash = hash
	if err := n.ValidateBlockCandidate(newBlock); err != nil {
		Validator.Stake -= 10
		return n.Blockchain, n.BlockchainHead, err
	} else {
		n.Blockchain.Blocks = append(n.Blockchain.Blocks, newBlock)
	}
	return n.Blockchain, blockBytes, nil

}

func (n LLxPosNetwork) ValidateChain() error {
	return nil
}

func (n LLxPosNetwork) ValidateBlockCandidate(newBlock *primitives.Block) error {

	return nil
}

func (n LLxPosNetwork) NewNode(addr string, stake uint) []*Node {
	newNode := &Node{
		Stake:   stake,
		Account: addr,
	}
	n.Validators = append(n.Validators, newNode)
	return n.Validators
}

func (n LLxPosNetwork) SelectValidator() (*Node, error) {
	var validatorPool []*Node
	totalStake := uint(0)
	for _, node := range n.Validators {
		if node.Stake > 0 {
			validatorPool = append(validatorPool, node)
			totalStake += node.Stake
		}
	}
	if validatorPool == nil {
		return nil, errors.New("there are no nodes with stake in the network")
	}

	validatorNumber := rand.Intn(int(totalStake))
	tmp := uint(0)
	for _, node := range n.Validators {
		tmp += node.Stake
		if validatorNumber < int(tmp) {
			return node, nil
		}

	}
	return nil, errors.New("a validator should have been picked but wasn't")

}

var Network LLxPosNetwork

func init() {

	Network.NewNode("Llx746869732069732061207368697474792073656e74656e636520746f2074657374Lx41444452", 100)

}
