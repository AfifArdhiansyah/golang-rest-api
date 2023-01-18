package models

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"type:bigint;primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
