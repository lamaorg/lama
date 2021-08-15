package primitives

import "math/big"

type iState interface {
	GetCurrentState()
	SetCurrentState()
	ValidateState()
}

type State struct {
	ID         *big.Int
	Time       uint64
	snapShot   *SnapShot
	SnapShotID *big.Int
}

type SnapShot struct {
	ID               *big.Int
	SecPubKey        []byte
	Genesis          iBlock
	ChainID          *big.Int
	Time             uint64
	Receipts         []Receipts
	NumTx            uint32
	NumWallets       uint32
	LastHash         []byte
	LastTxID         *big.Int
	ChainValue       *big.Int
	NumAccounts      *big.Int
	LastCodeExecuted string
	LastProof        string
}

type Receipts struct {
	ID               *big.Int
	TxID             *big.Int
	Time             uint64
	Recipient        string
	From             string
	Value            *big.Int
	Fee              *big.Int
	InBlockID        *big.Int
	ValidatorAddress string
	ValidatorHash    string
	Signature        []byte
}

type SnapShots []*SnapShot
type InvalidSnapShots []*big.Int
type ValidatedSnapShots []*big.Int
