package models

type Content struct {
	Reference int
	Audience  []User
	Meta      Meta
	Privacy   []string
	Flags     []string
}
