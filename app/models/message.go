package models

import (
	"time"
)

type Message struct {
	ID        int64     `gorm:"type:bigint;primary_key;auto_increment" json:"id"`
	Message   string    `gorm:"type:varchar(255)" json:"message"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
