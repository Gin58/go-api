package entity

import "time"

type Product struct {
	Id uint
	Title string `gorm:"size:128"`
	Content string `gorm:"size:256"`
	CreatedAt time.Time
}