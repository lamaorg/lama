package api

/*
	This file is used to bootstrap the lama blockchain
*/

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"github.com/lamaorg/lama/common"
	"github.com/lamaorg/lama/internals/primitives"

	"io/ioutil"

	"os"
)

func Bootstrapper() {

	// generate chain keys
	var sc common.SecureKeys
	k := sc.Create()

	var keys bytes.Buffer
	enc := json.NewEncoder(&keys)

	err := enc.Encode(&k)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll("./.keys", 0777)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./.keys/secure.key", keys.Bytes(), 0777)
	if err != nil {
		panic(err)
	}

	primitives.NewBlockchain()

}

func init() {
	Bootstrapper()

}

func generatePersonalSecretForWallet() ([]byte, error) {

	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil

}
