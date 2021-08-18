package Transaction

import (
	"github.com/lamaorg/lama/internals/primitives"
	"math/big"
)

var TXTYPE = TxType

type Base struct {
	Version   uint8    `json:"version"`
	Type      string   `json:"tx_type"`
	LockTime  int64    `json:"lock_time"`
	Fee       *big.Int `json:"fee"`
	From      string
	To        string
	Amount    *big.Int
	Memo      string
	Matures   int64
	Expires   int64
	Signature []byte
}

type Transaction struct {
	Base
}

type CoinbaseTx interface {
	IsCoinbase() bool
	GetCoinInfo() *primitives.Coin
}

func (t *Transaction) CoinbaseType() *Transaction {

	return t
}
