package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      string `gorm:"default:user"`
}

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	Stock       int
}
