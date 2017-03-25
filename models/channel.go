package models

type Channel struct {
	Model       *Model
	Acl         *Acl
	Name        string
	Description string
	Meta        string
	Namespace   string
	Privacy     string
	Rules       string
}

func (c *Channel) Save() bool {
	//convert to json and save
	return true
}

func (c *Channel) Destroy() bool {
	//Soft delete record
	return true
}
