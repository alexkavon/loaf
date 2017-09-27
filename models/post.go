package models

type Post struct {
	Loaf         string
	Audience     []User
	Contributors []User
	Meta         Meta
	Privacy      []string
	Parent       string
	Content      []Content
	Flags        []string
}
