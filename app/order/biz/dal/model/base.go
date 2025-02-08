package model

import "time"

type Base struct {
	Id        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
