package storage

import (
	badger "github.com/dgraph-io/badger/v3"
)

const localDbPath = "./tmp/lamallx"

type LocalStorage struct {
	Database *badger.DB
}

func (l *LocalStorage) Open(path string) {

	var err error
	l.Database, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		panic(err)
	}

}

func (l *LocalStorage) StoreKeyValue(key []byte, value []byte, path string) {

	new(LocalStorage).Open(path)
	db := l.Database
	defer db.Close()

	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})

	if err != nil {
		panic(err)
	}

}
