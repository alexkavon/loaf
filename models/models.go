package models

import (
	"time"
	//"encoding/json"
)

type Model struct {
	Id      int
	Created time.Time
	Updated time.Time
	Deleted *time.Time
}
