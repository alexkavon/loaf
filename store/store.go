package store

import (
	"github.com/boltdb/bolt"
	"gitlab.com/alexkavon/loaf/store"
)

type Store struct {
	Name string
	Host string
	Port int64
}

func NewStore() Store {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return
	}
	defer db.Close()
}
