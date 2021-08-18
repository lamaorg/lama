package api

/*
	This file is used to bootstrap the lama blockchain
*/

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

func Bootstrapper() {

	var api API
	// generate chain keys
	k := api.Keys.Create()

	var keys bytes.Buffer
	enc := gob.NewEncoder(&keys)

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
