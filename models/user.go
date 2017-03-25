package models

type User struct {
	Model         *Model
	Email         string
	Phone         string
	Settings      *Settings
	Status        string
	Subscriptions []Channel
	Level         int
}

func (u *User) Save() bool {
	//Save User
	return true
}

func (u *User) Destory() bool {
	//Cascade over user, single owner Channels, Settings, Conversations, etc.
	return true
}
