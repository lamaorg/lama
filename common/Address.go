package common

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
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
	hash := sha256.New()
	var h HexOperator
	hash.Write(xy[0])
	hash.Write(xy[1])
	hh := hash.Sum(nil)

	return h.EncodeHex(hh, "ADDRESS")

}

type Address struct{ raw *RawAddress }

func (a Address) New() Address {

	r := new(RawAddress)
	r.generateKeys()
	a.raw = r
	return a

}

func (a Address) GetPubX() []byte {

	var h HexOperator
	d := h.DecodeHex(a.raw.LlxAddress)
	return d.ToBytes

}

func (a Address) Serialize() []byte {
	addr, _ := json.Marshal(a)
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
