package store

import (
    "encoding/json"
	"github.com/boltdb/bolt"
)

type StorePlugin interface {
	NewStore() (*Store, error)
	Save() error
	Delete() error
	Prepare() error
}

type Store struct {
    Name string
    db *bolt.DB
}

func (s *StorePlugin) NewStore() (*Store, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
        return nil, err
	}
	defer db.Close()

    tx, err := db.Begin(true)
    if err != nil {
        return nil, err
    }

    defer tx.Rollback()

    if _, err := tx.CreateBucketIfNotExists([]byte("loafs")); err != nil {
        return nil, err
    }

    tx.Commit()

	return &Store{Name: "bolt", db: db}, nil
}

func (s *Store) Save() error {
    db := s.db
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    data, err := json.Marshall()
    return nil
}
