package models

import "time"

type MasjidTest struct {
	ID            uint      `json:"Id" gorm:"primaryKey;column:ID" `
	MasjidName    string    `json:"masjidName" gorm:"column:MASJID_NAME"`
	MasjidAddress string    `json:"masjidAddress" gorm:"column:MASJID_ADRESS"`
	City          string    `json:"city" gorm:"column:CITY"`
	Province      string    `json:"province" gorm:"column:PROVINCE"`
	BkmPIC        string    `json:"bkmpic" gorm:"column:BKM_PIC"`
	BkmPhone      string    `json:"bkmPhone" gorm:"column:BKM_PHONE"`
	CreatedBy     string    `json:"createdBy" gorm:"column:CREATED_BY"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:CreatedAt"`
	ModifiedBy    string    `json:"modifiedBy" gorm:"column:MODIFIED_BY"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:UpdatedAt"`
}
