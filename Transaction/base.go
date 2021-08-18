package Transaction

import (
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
	GetCoinInfo() string
}

func (t *Transaction) CoinbaseType() *Transaction {

	return t
}
