package Transaction

import (
	"github.com/lamaorg/lama/common/address"
	"github.com/lamaorg/lama/internals/primitives"
	"math/big"
	"time"
)

type Coinbase struct {
	ID                  *big.Int
	Currency            *primitives.Currency
	CoinCache           map[string]string
	Address             string
	Issuer              string
	Owner               string
	Minter              string
	Burner              string
	Operator            string
	MaxIssuance         *big.Int
	PremintIssuance     *big.Int
	MinIssuance         *big.Int
	InflationPerAnnum   *big.Int
	CashbackEnable      bool
	CashbackValuePerOne *big.Int
	MaxDivider          int
	Pairing             map[string]*PairCoin
}

type PairCoin struct {
	ID              string
	IsInternal      bool
	IsExternal      bool
	Name1           string
	Name2           string
	Value1          *big.Int
	Value2          *big.Int
	Liquidity1      *big.Int
	Liquidity2      *big.Int
	HistoricalValue map[int64]*big.Int
	TimeStamp       int64
}

type iCoinbase interface {
	GenesisTx()
	AssignCurrency()
	IsIssuer() bool
	IsOwner() bool
	IsMinter() bool
	IsBurner() bool
	IsOperator() bool
	GetNumMinted() int
	GetNumAvailable() int
	GetNumUnminted() int
	GetMaxIssuance() *big.Int
	GetMinIssuance() *big.Int
	CalcCurrentInflation() *big.Int
	GetPair(name1, name2 string) *PairCoin
	SetPair(args ...string)
	Mint(qty int32) (bool, error)
	Burn(qty int32) (bool, error)
}

func CreateNewInternalCoinbase() *Coinbase {
	cb := new(Coinbase)
	cb.ID = new(big.Int).SetUint64(uint64(time.Now().UnixNano()))
	cb.CoinCache = make(map[string]string, 0)
	_, buf := address.GenerateNewAddress()
	cb.Address = "LLx" + buf

	return cb

}
