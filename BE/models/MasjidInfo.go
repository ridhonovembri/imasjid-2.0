package models

import "time"

// type MasjidInfo struct {
// 	ID            uint      `json:"id" gorm:"primaryKey;column:ID"`
// 	MasjidName    string    `json:"masjidName" gorm:"column:MASJID_NAME"`
// 	MasjidAddress string    `json:"masjidAddress" gorm:"column:MASJID_ADDRESS"`
// 	City          string    `json:"city" gorm:"column:CITY"`
// 	Province      string    `json:"province" gorm:"PROVINCE"`
// 	BkmPIC        string    `json:"bkmpic" gorm:"BKM_PIC"`
// 	BkmPhone      string    `json:"bkmPhone" gorm:"BKM_PHONE"`
// 	CreatedBy     string    `json:"createdBy" gorm:"CREATED_BY"`
// 	CreatedAt     time.Time `json:"createdAt" gorm:"CreatedAt"`
// 	ModifiedBy    string    `json:"modifiedBy" gorm:"MODIFIED_BY"`
// 	UpdatedAt     time.Time `json:"updatedAt" gorm:"UpdatedAt"`
// }

// type Book struct {
// 	Id          int64  `gorm:"primaryKey" json:"id"`
// 	Title       string `gorm:"type:varchar(300)" json:"title"`
// 	Description string `gorm:"type:text" json:"description"`
// 	Author      string `gorm:"type:varchar(300)" json:"author"`
// 	PublishDate string `gorm:"type:date" json:"publish_date"`
// }

type MasjidInfo struct {
	ID            uint      `json:"Id" gorm:"primaryKey;column:ID"`
	MasjidName    string    `json:"MasjidName" gorm:"column:MASJID_NAME"`
	MasjidAddress string    `json:"MasjidAddress" gorm:"column:MASJID_ADDRESS"`
	City          string    `json:"City" gorm:"column:CITY"`
	Province      string    `json:"Province" gorm:"column:PROVINCE"`
	BkmPIC        string    `json:"BkmPic" gorm:"column:BKM_PIC"`
	BkmPhone      string    `json:"BkmPhone" gorm:"column:BKM_PHONE"`
	CreatedBy     string    `json:"CreatedBy" gorm:"column:CREATED_BY"`
	CreatedAt     time.Time `json:"CreatedAt" gorm:"column:CREATED_AT"`
	ModifiedBy    string    `json:"ModifiedBy" gorm:"column:MODIFIED_BY"`
	UpdatedAt     time.Time `json:"UpdatedAt" gorm:"column:UPDATED_AT"`
}
