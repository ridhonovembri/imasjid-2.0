package models

import "time"

type MasjidInfo struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	MasjidName    string    `json:"masjidName"`
	MasjidAddress string    `json:"masjidAddress"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	BkmPIC        string    `json:"bkmpic"`
	BkmPhone      string    `json:"bkmPhone"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedByName string    `json:"createdByName"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedByName string    `json:"updatedByName"`
}
