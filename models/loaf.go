package models

import (
	"gitlab.com/alexkavon/loaf/store"
)

type Loaf struct {
	Acl         *Acl            `json:"acl"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Meta        Meta            `json:"meta"`
	Namespace   string          `json:"namespace"`
	Privacy     string          `json:"privacy"`
	Content     map[int]Content `json:"content"`
}

func (l *Loaf) Index() {}

func (l *Loaf) Store() {}

func (l *Loaf) Update() {}

func (l *Loaf) Destroy() {}
