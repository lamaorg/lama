package primitives

import (
	"github.com/lamaorg/lama/common"
	"math/big"
)

var TxType map[string]uint = make(map[string]uint, 0)

func init() {
	TxType["COINBASE"] = 0x0000
	TxType["WALLET"] = 0x1010
	TxType["CONTENT"] = 0x0201
	TxType["MINT"] = 0x0301
	TxType["BURN"] = 0x0302
	TxType["TRANSFER"] = 0x0303
	TxType["ALLOW"] = 0x0401
	TxType["SECURITY"] = 0x0555
	TxType["SYSTEM"] = 0x0999
	TxType["EVENT"] = 0x0606
	TxType["REVERT"] = 0x0990
	TxType["TOKEN"] = 0x0001
	TxType["VM"] = 0x7777
	TxType["CONTRACT"] = 0x7771
	TxType["SWAP"] = 0x7772
	TxType["EXCHANGE"] = 0x7773
	TxType["BUY"] = 0x7774
	TxType["REWARD"] = 0x7775

}

type Transaction struct {
	ID             string
	Nonce          string
	Proof          string
	Hash           string
	Value          *big.Int
	Signature      string
	NumInputs      uint32
	NumOutputs     uint32
	Time           uint64
	Inputs         map[int]*Input
	Outputs        map[int]*Output
	Validations    uint64
	Validated      bool
	Validator      common.Address
	ValidatorValue uint64
}

type Transactions []*Transaction

type Input struct {
	TxType      uint           `json:"tx_type"`
	From        common.Address `json:"tx_from"`
	To          common.Address `json:"tx_to"`
	IsPayable   bool
	IsAnonymous bool
	Data        []byte
	Time        uint64
	Amount      *big.Int
	TxFee       *big.Int
	State       string
}

type Output struct {
	HexInput string
	To       common.Address
	Spent    bool
	Amount   *big.Int
	Data     []byte
	State    string
}

func NewTransaction() {}
