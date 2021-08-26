package common

import (
	"encoding/hex"
	"github.com/keybase/go-keychain"
	"io/ioutil"
)

type LLamaKeys struct {
	Keychain keychain.Keychain
}

const KEYCHAINNAME = "LlamaKeysSet"

var LK *LLamaKeys

func init() {
	keychainKey, _ := ioutil.ReadFile("./.keys/secure.key")

	k, _ := keychain.NewKeychain(KEYCHAINNAME, hex.EncodeToString(keychainKey))
	LK = new(LLamaKeys)
	LK.Keychain = k

}

func (L *LLamaKeys) AddKey(serviceName string, keyName string, keyLabel string, keyPass []byte, keyGroup string) {

	item := keychain.NewGenericPassword(serviceName, keyName, keyLabel, keyPass, keyGroup)
	item.SetMatchLimit(keychain.MatchLimitOne)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	item.UseKeychain(LK.Keychain)
	err := keychain.AddItem(item)
	if err != nil {
		panic(err)
	}

}

func (L *LLamaKeys) GetAllAccountsForService(serviceName string) ([]string, error) {

	accounts, err := keychain.GetAccountsForService(serviceName)
	if err != nil {
		return nil, err
	}
	return accounts, nil

}
