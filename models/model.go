package models

type Model interface {
	All()
	One()
	Validate()
	Store()
	Update()
	Destroy()
}
