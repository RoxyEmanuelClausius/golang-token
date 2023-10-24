package models

import "time"

type customer struct {
	Id               int       `gorm:"primaryKey" json:"id"`
	full_name        string    `gorm:"type:varchar(250)" json:"track_name"`
	username         string    `gorm:"type:varchar(100)" json:"username"`
	email            string    `gorm:"type:varchar(150);unique_index" json:"email"`
	password         string    `gorm:"type:varchar(250)" json:"password"`
	fpass_code       string    `gorm:"type:varchar(250)" json:"fpass_code"`
	fpass_expired    time.Time `gorm:"type:date" json:"fpass_expired"`
	oauth_token1     string    `gorm:"type:text('medium')" json:"oauth_token1"`
	oauth_token2     string    `gorm:"type:text('medium')" json:"oauth_token2"`
	oauth_token3     string    `gorm:"type:text('medium')" json:"oauth_token3"`
	alternatif_link1 string    `gorm:"type:varchar(250)" json:"alternatif_link1"`
	alternatif_link2 string    `gorm:"type:varchar(250)" json:"alternatif_link2"`
	alternatif_link3 string    `gorm:"type:varchar(250)" json:"alternatif_link3"`
	jwt_token        string    `gorm:"type:text" json:"jwt_token"`
	created_at       time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	updated_at       time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`
	deleted_at       time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"deleted_at"`
}
