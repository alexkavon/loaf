package models

type Conversation struct {
	Model        *Model
	Channel      string
	Audience     []User
	Contributors []User
	Meta         string
	Privacy      []string
	Parent       string
	Content      string
	Flags        []string
}
