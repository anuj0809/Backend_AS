package models

// create the model
type Players struct {
	ID      uint64 `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"size:15"`
	Country string `json:"country" gorm:"size:2"`
	Score   int    `json:"score"`
}
