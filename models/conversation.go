package models

type Conversation struct {
	Loaf         string
	Audience     []User
	Contributors []User
	Meta         Meta
	Privacy      []string
	Parent       string
	Content      []Content
	Flags        []string
}
