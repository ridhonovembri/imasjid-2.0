package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedByName string    `json:"createdByName"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedByName string    `json:"updatedByName"`
}
