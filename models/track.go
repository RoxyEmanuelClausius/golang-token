package models

type track struct {
	Id             int64   `gorm:"primaryKey" json:"id"`
	track_name     string  `gorm:"type:varchar(250)" json:"track_name"`
	track_type     string  `gorm:"type:varchar(250)" json:"track_type"`
	track_status   string  `gorm:"type:varchar(250)" json:"track_status"`
	track_durasion string  `gorm:"type:varchar(250)" json:"track_durasion"`
	track_bpm      string  `gorm:"type:varchar(250)" json:"track_bpm"`
	track_key      string  `gorm:"type:varchar(50)" json:"track_key"`
	track_artwork  string  `gorm:"type:varchar(250)" json:"track_artwork"`
	track_file1    string  `gorm:"type:varchar(250)" json:"track_file1"`
	track_file2    string  `gorm:"type:varchar(250)" json:"track_file2"`
	track_file3    string  `gorm:"type:varchar(250)" json:"track_file3"`
	track_price1   float32 `gorm:"type:decimal(15,0)" json:"track_price1"`
	track_price2   float32 `gorm:"type:decimal(15,0)" json:"track_price2"`
	track_price3   float32 `gorm:"type:decimal(15,0)" json:"track_price3"`
}
