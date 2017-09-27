package models

type Model interface {
	Index()
	Store()
	Update()
	Destroy()
}
