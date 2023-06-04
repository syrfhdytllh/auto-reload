package models

import "time"

type Status struct {
	Id        uint       `gorm:"primaryKey"`
	Water     string     `gorm:"varchar(100)" json:"water"`
	Air       string     `gorm:"varchar(100)" json:"air"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
