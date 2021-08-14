package primitives

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	lru "github.com/hashicorp/golang-lru"
	"github.com/lamaorg/lama/common"
	"golang.org/x/crypto/sha3"
	"math/big"
)

const HEXTYPE = "INTERNAL"

var EmptyRootHash = []byte("fe127cd046bf452c49ad2e659dd49650b7e672a0f8348f819fab3baed47827b8")

type Header struct {
	ParentHash      []byte   `json:"parentHash"`
	Coinbase        []byte   `json:"coinbase_operator"`
	StateRootHash   []byte   `json:"state_root"`
	TxRootHash      []byte   `json:"tx_root"`
	ProofRootHash   []byte   `json:"proof_root"`
	ReceiptRootHash []byte   `json:"receipt_root"`
	Time            uint64   `json:"timestamp"`
	FeeLimit        uint64   `json:"fee_limit"`
	TotalFees       uint64   `json:"total_fees"`
	ExtraData       []byte   `json:"extra_data"`
	BigNumber       *big.Int `json:"unique_big_number"`
	ProofTargetHash []byte   `json:"proof_target"`
	BaseFee         *big.Int `json:"base_fee"`
}

type serializedHeader struct {
	Number    string
	FeeLimit  string
	TotalFees string
	Time      string
	Extra     string
	BaseFee   string
	Hash      string
}

func (h Header) serNumber() string {

	return common.EncodeBig(h.BigNumber)

}

func (h Header) serialize(ser *serializedHeader) *serializedHeader {

	fl := common.EncodeUint64(h.FeeLimit)
	tf := common.EncodeUint64(h.TotalFees)
	t := common.EncodeUint64(h.Time)
	ser._serUint64(fl, tf, t)
	ser.BaseFee = common.EncodeBig(h.BaseFee)
	ser.Number = h.serNumber()
	ser.Extra = common.EncodeNewHex(h.ExtraData, HEXTYPE)
	return ser

}

func (ser *serializedHeader) _serUint64(feelimit, totalfee, time string) {
	ser.FeeLimit = feelimit
	ser.TotalFees = totalfee
	ser.Time = time

}

func (h Header) Hash() []byte {
	m, _ := json.Marshal(h.serialize(new(serializedHeader)))
	var bufbytes []byte
	hash := sha3.New256()
	hash.Write(h.ParentHash)
	copy(bufbytes, hash.Sum(nil))
	hash.Write(m)
	hash2 := sha256.New()
	hash2.Write(bufbytes[:])
	b := make([]byte, hash.Size()+hash2.Size())
	b = append(b, hash.Sum(nil)...)
	b = append(b, hash2.Sum(nil)...)

	return hash2.Sum(b)
}

func (h *Header) EmptyBody() bool {
	return bytes.Compare(h.TxRootHash, EmptyRootHash) == 0
}

func (h *Header) EmptyReceipts() bool {
	return bytes.Compare(h.ReceiptRootHash, EmptyRootHash) == 0
}

type Body struct {
	Txs   []*Transaction
	Heads []*Header
}

type Block struct {
	header       *Header
	parentHeads  []*Header
	transactions Transactions
	hash         lru.Cache
	size         lru.Cache
	proof        lru.Cache
}

type externalBlock struct {
	Header *Header
	Txs    []*Transaction
	Heads  []*Header
}
