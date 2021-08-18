package api

/*
	This file is used to bootstrap the lama blockchain
*/

import (
	"bytes"
	"crypto/rand"

	"encoding/json"

	"io/ioutil"

	"os"
)

func Bootstrapper() {

	var api API
	// generate chain keys
	k := api.Keys.Create()

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
