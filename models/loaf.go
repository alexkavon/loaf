package models

type Loaf struct {
	Acl         *Acl   `json:"acl"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Meta        Meta   `json:"meta"`
	Namespace   string `json:"namespace"`
	Privacy     string `json:"privacy"`
	Rules       string `json:"rules"`
}

func (l *Loaf) Save() bool {
	//convert to json and save
	return true
}

func (l *Loaf) Destroy() bool {
	//Soft delete record
	return true
}
