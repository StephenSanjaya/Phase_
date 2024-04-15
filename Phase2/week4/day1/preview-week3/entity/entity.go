package entity

type Player struct {
	ID       int    `json:"player_id" gorm:"primaryKey;column:player_id"`
	Username string `json:"username" gorm:"not null"`
	TeamName string `json:"team_name" gorm:"not null"`
	Ranking  int    `json:"ranking" gorm:"not null"`
	Score    int    `json:"score" gorm:"not null"`
}
