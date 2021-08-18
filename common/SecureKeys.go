package common

import (
	"errors"
	d "github.com/kudelskisecurity/crystals-go/crystals-dilithium"
	k "github.com/kudelskisecurity/crystals-go/crystals-kyber"
)

type SecureKeys struct {
	kyberInst     *k.Kyber
	dilithiumInst *d.Dilithium
	shared        []byte
	seed          []byte
	//packed private key
	PSK []byte
	//packed public key
	PPK []byte
	DK1 []byte
	DK2 []byte
}

type SecureMessage struct {
	c            []byte
	sharedSecret []byte
	senderKey    []byte
	receiverKey  []byte
}

func (s *SecureKeys) Create() *SecureKeys {

	key := k.NewKyber512()
	d3 := d.NewDilithium2()
	s.kyberInst = key
	s.dilithiumInst = d3
	seed, _ := GenRandomBytes(64)
	s.seed = seed
	pk, sk := key.KeyGen(seed)

	s.PSK = sk
	s.PPK = pk
	dk1, dk2 := d3.KeyGen(s.seed)
	s.DK1 = dk1
	s.DK2 = dk2

	return s

}

//SendSS will send the public key from A to B to ack a shared secret
func (s SecureKeys) SendSS(msg []byte) *SecureMessage {

	c, sharedSecret := s.kyberInst.Encaps(s.PPK, msg)
	sm := new(SecureMessage)
	sm.c = c
	sm.sharedSecret = sharedSecret
	sm.senderKey = s.PPK

	return sm

}

//VerifySS check the msg from A and gets the shared secret from it
func (s SecureKeys) VerifySS(msg []byte) []byte {
	ss := s.kyberInst.Decaps(s.PSK, msg)
	return ss
}

func (s SecureKeys) Decrypt(sm SecureMessage) []byte {

	decrypted := s.kyberInst.Decrypt(s.PSK, sm.c)
	return decrypted

}

func (s SecureKeys) Encrypt(msg []byte) ([]byte, error) {
	seed, _ := GenRandomBytes(32)
	encrypted := s.kyberInst.Encrypt(s.PPK, msg, seed)
	if encrypted != nil {
		return encrypted, nil
	}
	return nil, errors.New("there was an error encrypting using the kyber encryption")

}

func (s SecureKeys) SecureSign(msg []byte) []byte {

	return s.dilithiumInst.Sign(s.DK2, msg)

}

func (s SecureKeys) SecureVerify(msg []byte, pub []byte, sig []byte) (bool, error) {

	v := s.dilithiumInst.Verify(pub, sig, msg)
	if v == true {
		return true, nil
	}
	return false, errors.New("invalid signature (dilithium)")
}
