package models

import (
	"time"
)

type Album struct {
	ID        int64     `gorm:"type:bigint;primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Artist    string    `gorm:"type:varchar(255)" json:"artist"`
	Price     int64     `gorm:"type:bigint" json:"price"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
