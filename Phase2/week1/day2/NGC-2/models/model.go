package models

type Heroes struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"imgUrl"`
}

type Villain struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	ImageURL string `json:"imgUrl"`
}
