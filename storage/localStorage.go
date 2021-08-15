package storage

import (
	"github.com/dgraph-io/badger"
)

const localDbPath = "./tmp/lamallx"

type LocalStorage struct {
	Database *badger.DB
}
