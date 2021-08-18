package address

import (
	"encoding/json"
	"github.com/lamaorg/lama/common"
	"golang.org/x/crypto/blake2b"

	"bytes"
	"github.com/lunixbochs/struc"
	"time"
)

type address struct {
	Keys      *common.SecureKeys
	Signature []byte
	Purpose   int
	Timestamp int64
	Hash      []byte
	MetaData  map[string]string
}

type AddressUnpack struct {
	Pub string `struc:[64]int8`
}

func (a *address) GetNew() *address {

	sc := new(common.SecureKeys)
	a.Keys = sc.Create()
	// signature will be based on the dilithium3 2 signed key (private key)
	a.Signature = a.Keys.SecureSign(a.Keys.DK1)
	a.Purpose = 1
	a.Timestamp = time.Now().UnixNano()
	h, _ := json.Marshal(a)
	hash := blake2b.Sum256(h)
	a.Hash = hash[:]
	return a

}

func (u *AddressUnpack) Pack() []byte {

	var buf bytes.Buffer
	err := struc.Pack(&buf, u)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()

}

func Validate(A *address, dk1 []byte, sig []byte) (bool, error) {

	var sc common.SecureKeys
	return sc.SecureVerify(A.Keys.DK1, dk1, sig)

}

func GenerateNewAddress() (*address, string) {
	a := new(address)
	addr := a.GetNew()
	unpacked := new(AddressUnpack)

	unpacked.Pub = common.Encode(addr.Hash)

	//spew.Dump(unpacked.Pack())
	var buf bytes.Buffer
	_ = struc.Pack(&buf, unpacked)

	return addr, buf.String()
}
