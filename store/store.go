package store

import (
	"github.com/boltdb/bolt"
)

type Store {
	Name	string
	Config	*StoreConfig
}

type StoreConfig struct {
	Host	string
	Port	int64
}

type StorePlugin interface {
	NewStore() Store
	Save() error
	Delete() error
	Prepare() Transaction
}

func NewStore() Store {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return
	}
	defer db.Close()

	return &Store{Name: "Bolt", Config: &StoreConfig{Host: "my.db", Port: 8080}}
}
