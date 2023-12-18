package models

import (
	"time"
)

// type Book struct {
// 	Id          int64  `gorm:"primaryKey" json:"id"`
// 	Title       string `gorm:"type:varchar(300)" json:"title"`
// 	Description string `gorm:"type:text" json:"description"`
// 	Author      string `gorm:"type:varchar(300)" json:"author"`
// 	PublishDate string `gorm:"type:date" json:"publish_date"`
// }

type MasjidConfig struct {
	ID                     int8       `json:"Id" gorm:"primaryKey;column:ID"`
	MinutesToAdzanShubuh   int       `json:"MinutesToAdzanShubuh" gorm:"column:MINUTES_TO_ADZAN_SHUBUH"`
	MinutesToAdzanDzuhur   int       `json:"MinutesToAdzanDzuhur" gorm:"column:MINUTES_TO_ADZAN_DZUHUR"`
	MinutesToAdzanAshar    int       `json:"MinutesToAdzanAshar" gorm:"column:MINUTES_TO_ADZAN_ASHAR"`
	MinutesToAdzanMaghrib  int       `json:"MinutesToAdzanMaghrib" gorm:"column:MINUTES_TO_ADZAN_MAGHRIB"`
	MinutesToAdzanIsya     int       `json:"MinutesToAdzanIsya" gorm:"column:MINUTES_TO_ADZAN_ISYA"`
	MinutesToIqomahShubuh  int       `json:"MinutesToIqomahShubuh" gorm:"column:MINUTES_TO_IQOMAH_SHUBUH"`
	MinutesToIqomahDzuhur  int       `json:"MinutesToIqomahDzuhur" gorm:"column:MINUTES_TO_IQOMAH_DZUHUR"`
	MinutesToIqomahAshar   int       `json:"MinutesToIqomahAshar" gorm:"column:MINUTES_TO_IQOMAH_ASHAR"`
	MinutesToIqomahMaghrib int       `json:"MinutesToIqomahMaghrib" gorm:"column:MINUTES_TO_IQOMAH_MAGHRIB"`
	MinutesToIqomahIsya    int       `json:"MinutesToIqomahIsya" gorm:"column:MINUTES_TO_IQOMAH_ISYA"`
	MinutesToAdzanPrep     int       `json:"MinutesToAdzanPrep" gorm:"column:MINUTES_TO_ADZAN_PREP"`
	LabelShubuh            string    `json:"LabelShubuh" gorm:"column:LABEL_SHUBUH"`
	LabelSyuruq            string    `json:"LabelSyuruq" gorm:"column:LABEL_SYURUQ"`
	LabelDzuhur            string    `json:"LabelDzuhur" gorm:"column:LABEL_DZUHUR"`
	LabelAshar             string    `json:"LabelAshar" gorm:"column:LABEL_ASHAR"`
	LabelMaghrib           string    `json:"LabelMaghrib" gorm:"column:LABEL_MAGHRIB"`
	LabelIsya              string    `json:"LabelIsya" gorm:"column:LABEL_ISYA"`
	AdjustShubuh           int       `json:"AdjustShubuh" gorm:"column:ADJUST_SHUBUH"`
	AdjustSyuruq           int       `json:"AdjustSyuruq" gorm:"column:ADJUST_SYURUQ"`
	Adjustzuhur            int       `json:"AdjustDzuhur" gorm:"column:ADJUST_DZUHUR"`
	AdjustAshar            int       `json:"AdjustAshar" gorm:"column:ADJUST_ASHAR"`
	AdjustMaghrib          int       `json:"AdjustMaghrib" gorm:"column:ADJUST_MAGHRIB"`
	AdjustIsya             int       `json:"AdjustIsya" gorm:"column:ADJUST_ISYA"`
	TextBeforeAdzan        string    `json:"TextBeforeAdzan" gorm:"column:TEXT_BEFORE_ADZAN"`
	TextBeforeIqomah       string    `json:"TextBeforeIqomah" gorm:"column:TEXT_BEFORE_IQOMAH"`
	TextDuringSholat1      string    `json:"TextDuringSholat1" gorm:"column:TEXT_DURING_SHOLAT1"`
	TextDuringSholat2      string    `json:"TextDuringSholat2" gorm:"column:TEXT_DURING_SHOLAT2"`
	Latitude               string    `json:"Latitude" gorm:"column:LATITUDE"`
	Longitude              string    `json:"Longitude" gorm:"column:LONGITUDE"`
	BgImgSholat            string    `json:"BgImgSholat" gorm:"column:BG_IMG_SHOLAT"`
	IconDuringSholat       string    `json:"IconDuringSholat" gorm:"column:ICON_DURING_SHOLAT"`
	IntervalSlide          int       `json:"IntervalSlide" gorm:"column:INTERVAL_SLIDE"`
	IntervalHadist         int       `json:"IntervalHadist" gorm:"column:INTERVAL_HADIST"`
	SholatDuration         int       `json:"SholatDuration" gorm:"column:SHOLAT_DURATION"`
	SoundLocation          string    `json:"SoundLocation" gorm:"column:SOUND_LOCATION"`
	FontSizeMarquee        int       `json:"FontSizeMarquee" gorm:"column:FONTSIZE_MARQUEE"`
	FontSizeHadist         int       `json:"FontSizeHadist" gorm:"column:FONTSIZE_HADIST"`
	FontSizeMasjidname     int       `json:"FontSizeMasjidname" gorm:"column:FONTSIZE_MASJIDNAME"`
	FontSizeMasjidaddress  int       `json:"FontSizeMasjidaddress" gorm:"column:FONTSIZE_MASJIDADDRESS"`
	FontSizeDate           int       `json:"FontSizeDate" gorm:"column:FONTSIZE_DATE"`
	FontSizeClock          int       `json:"FontSizeClock" gorm:"column:FONTSIZE_CLOCK"`
	FontSizePrayertime     int       `json:"FontSizePrayertime" gorm:"column:FONTSIZE_PRAYERTIME"`
	FontSizePrayerlabel    int       `json:"FontSizePrayerlabel" gorm:"column:FONTSIZE_PRAYERLABEL"`
	CreatedBy              string    `json:"CreatedBy" gorm:"column:CREATED_BY"`
	CreatedAt              time.Time `json:"CreatedAt" gorm:"column:CREATED_AT"`
	ModifiedBy             string    `json:"ModifiedBy" gorm:"column:MODIFIED_BY"`
	UpdatedAt              time.Time `json:"UpdatedAt" gorm:"column:UPDATED_AT"`
}
