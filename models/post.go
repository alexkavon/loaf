package models

type Post struct {
	Loaf         string
	Audience     []Loaf
	Contributors []Loaf
	Meta         Meta
	Privacy      []string
	Parent       string
	Content      []Content
	Flags        []string
}
