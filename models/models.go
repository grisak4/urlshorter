package models

import "time"

type ShortUrl struct {
	ID          uint      `gorm:"unique;primaryKey" json:"id"`
	Url         string    `gorm:"unique" json:"url"`
	ShortCode   string    `gorm:"unique" json:"shortCode"`
	AccessCount int       `json:"accessCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
