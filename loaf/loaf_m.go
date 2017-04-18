package loaf

import (
	"encoding/json"
	"time"
)

type loaf struct {
	Id      int       `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Name    string    `json:"name"`
	About   string    `json:"about"`
	Rules   string    `json:"rules"`
	Acl     []string  `json:"-"`
}

func newLoaf(data string) *loaf {
	return &loaf{}
}

func updateLoaf(data string) *loaf {
	return &loaf{}
}

func deleteLoaf() bool {
	return true
}
