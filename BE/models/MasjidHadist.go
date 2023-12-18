package models

import "time"

// type Book struct {
// 	Id          int64  `gorm:"primaryKey" json:"id"`
// 	Title       string `gorm:"type:varchar(300)" json:"title"`
// 	Description string `gorm:"type:text" json:"description"`
// 	Author      string `gorm:"type:varchar(300)" json:"author"`
// 	PublishDate string `gorm:"type:date" json:"publish_date"`
// }

type MasjidHadist struct {
	ID            uint      `json:"Id" gorm:"primaryKey;column:ID"`
	HadistContent string    `json:"HadistContent" gorm:"column:HADIST_CONTENT"`
	HadistFrom    string    `json:"HadistFrom" gorm:"column:HADIST_FROM"`
	CreatedBy     string    `json:"CreatedBy" gorm:"column:CREATED_BY"`
	CreatedAt     time.Time `json:"CreatedAt" gorm:"column:CREATED_AT"`
	ModifiedBy    string    `json:"ModifiedBy" gorm:"column:MODIFIED_BY"`
	UpdatedAt     time.Time `json:"UpdatedAt" gorm:"column:UPDATED_AT"`
}
