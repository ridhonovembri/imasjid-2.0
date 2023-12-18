package models

import "time"

// type Book struct {
// 	Id          int64  `gorm:"primaryKey" json:"id"`
// 	Title       string `gorm:"type:varchar(300)" json:"title"`
// 	Description string `gorm:"type:text" json:"description"`
// 	Author      string `gorm:"type:varchar(300)" json:"author"`
// 	PublishDate string `gorm:"type:date" json:"publish_date"`
// }

type MasjidSlides struct {
	ID         uint      `json:"Id" gorm:"primaryKey;column:ID"`
	Sequence   string    `json:"Sequence" gorm:"column:SEQUENCE"`
	ImgName    string    `json:"ImgName" gorm:"column:IMG_NAME"`
	ImgDesc    string    `json:"ImgDesc" gorm:"column:IMG_DESC"`
	ImgSource  string    `json:"ImgSource" gorm:"column:IMG_SOURCE"`
	CreatedBy  string    `json:"CreatedBy" gorm:"column:CREATED_BY"`
	CreatedAt  time.Time `json:"CreatedAt" gorm:"column:CREATED_AT"`
	ModifiedBy string    `json:"ModifiedBy" gorm:"column:MODIFIED_BY"`
	UpdatedAt  time.Time `json:"UpdatedAt" gorm:"column:UPDATED_AT"`
}
