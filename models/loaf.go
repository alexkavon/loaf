package models

import (
//    "gitlab.com/alexkavon/loaf/store"
//    "gopkg.in/validator.v2"
)

type Loaf struct {
    Acl         *Acl            `json:"acl"`
    Name        string          `json:"name" validate:"min=1"`
    Description string          `json:"description"`
    Meta        Meta            `json:"meta"`
    Namespace   string          `json:"namespace"`
    Privacy     string          `json:"privacy"`
    Content     map[int]Content `json:"content"`
}

func NewLoaf(name string) *Loaf {
    return &Loaf{Name: name}
}

func (l *Loaf) Index() {}

func (l *Loaf) Store() {}

func (l *Loaf) Update() {}

func (l *Loaf) Destroy() {}
