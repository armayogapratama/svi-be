package model

import (
	"time"

	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title        string    `json:"title" gorm:"not null; required"`
	Content      string    `json:"content" gorm:"not null; required"`
	Category     string    `json:"category" gorm:"not null; required"`
	Status       string    `json:"status" gorm:"required; check:status in ('publish', 'draft', 'thrash')"`
	Created_Date time.Time `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_Date time.Time `json:"updated_date" gorm:"default:CURRENT_TIMESTAMP"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Search struct {
	Status string `json:"status"`
}
