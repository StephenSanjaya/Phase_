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

type CriminalReports struct {
	ID          int64  `json:"id"`
	HeroID      int64  `json:"hero_id"`
	VillainID   int64  `json:"villain_id"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

type SuccessMessage struct {
	Message string            `json:"message"`
	Status  int64             `json:"status"`
	Datas   []CriminalReports `json:"datas,omitempty"`
	Data    *CriminalReports  `json:"data,omitempty"`
}
