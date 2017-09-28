package models

type Content struct {
	Audience  []*Loaf
	Meta      Meta
	Privacy   []string
	Flags     []string
}
