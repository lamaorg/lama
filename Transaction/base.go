package Transaction

import "github.com/lamaorg/lama/internals/dnaProof"

type Base struct {
	Version          uint8  `json:"version"`
	Type             string `json:"tx_type"`
	LockTime         int64  `json:"lock_time"`
	Fee              uint64 `json:"fee"`
	SigPubKey        []byte `json:"sig_pub_key"`
	Sig              []byte `json:"sig"`
	Proof            *dnaProof.Proof
	Metadata         map[string]string
	sigPrivKey       []byte
	cachedHash       []byte
	cachedActualSize uint64
}
