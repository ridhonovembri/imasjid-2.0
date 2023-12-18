package models

import "time"

type MasjidMarquee struct {
	ID          uint      `json:"Id" gorm:"primaryKey;column:ID"`
	MarqueeText string    `json:"MarqueeText" gorm:"column:MARQUEE_TEXT"`
	CreatedBy   string    `json:"CreatedBy" gorm:"column:CREATED_BY"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"column:CREATED_AT"`
	ModifiedBy  string    `json:"ModifiedBy" gorm:"column:MODIFIED_BY"`
	UpdatedAt   time.Time `json:"UpdatedAt" gorm:"column:UPDATED_AT"`
}
