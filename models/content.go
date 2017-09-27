package models

type Content struct {
	Audience  []User
	Meta      Meta
	Privacy   []string
	Flags     []string
}
