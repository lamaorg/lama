package common

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"golang.org/x/crypto/ripemd160"
	"io/ioutil"
	"math/big"
)

type RawAddress struct {
	ECDSAPriv  *ecdsa.PrivateKey
	Pub        ecdsa.PublicKey
	Sig        []byte
	LlxAddress string
}

func (r *RawAddress) generateKeys() {

	key, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)

	r.ECDSAPriv = key
	r.Pub = key.PublicKey
	r.LlxAddress = r.ToString()
	r.Sig, _ = r.ECDSAPriv.Sign(rand.Reader, []byte(r.LlxAddress), crypto.SHA3_256)

}

func (r *RawAddress) sign() (*big.Int, *big.Int) {

	rr, s, _ := ecdsa.Sign(rand.Reader, r.ECDSAPriv, []byte(r.LlxAddress))
	return rr, s

}

func GenRandomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func (r *RawAddress) SignASN(data []byte) []byte {

	s, _ := ecdsa.SignASN1(rand.Reader, r.ECDSAPriv, data)
	return s

}

func VerifyPublicSignature(pub *ecdsa.PublicKey, hashed []byte, sig []byte) bool {
	return ecdsa.VerifyASN1(pub, hashed, sig)
}

func (r *RawAddress) GetPublic() crypto.PublicKey {
	return r.ECDSAPriv.Public()
}

func (r *RawAddress) ToString() string {

	x := EncodeBig(r.Pub.X)
	y := EncodeBig(r.Pub.Y)
	xy := [][]byte{[]byte(x), []byte(y)}
	xyj := bytes.Join(xy, []byte("_"))

	hash2 := Encode(xyj)
	h1 := ripemd160.New()
	h1.Write([]byte(hash2))
	var h HexOperator
	return h.EncodeHex(h1.Sum(nil), "ADDRESS")

}

type Address struct{ raw *RawAddress }

func (a Address) New() Address {

	r := new(RawAddress)
	r.generateKeys()
	a.raw = r
	ser := a.Serialize()
	StoreAddress(ser)
	return a

}

func (a Address) GetPubX() []byte {

	var h HexOperator
	d := h.DecodeHex(a.raw.LlxAddress)
	return d.ToBytes

}

func (a Address) Serialize() []byte {
	addr, _ := json.Marshal(a.raw)
	return addr
}

func UnserializeAddress(ab []byte) *Address {

	var A Address
	err := json.Unmarshal(ab, &A)
	if err != nil {
		return nil
	}
	return &A

}

func VerifyAddress(a Address) bool {

	return ecdsa.VerifyASN1(&a.raw.Pub, []byte(a.raw.LlxAddress), a.raw.Sig)

}

func StoreAddress(marshaled []byte) {

	err := ioutil.WriteFile("./wallet.json", []byte(marshaled), 0644)
	if err != nil {
		return
	}

}

func LoadWallet(address string, key string) {

}
