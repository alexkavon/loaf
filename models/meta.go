package models

type Meta struct {
	Id      int64     `json:"id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
