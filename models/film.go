package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Film struct {
	Title     string
	Director  string
	Year      int
	BoxOffice float64
}

type MovieDB struct {
	gorm.Model
	Title     string `gorm:"uniqueIndex:title_yr_idx"`
	Director  string
	Year      int `gorm:"uniqueIndex:title_yr_idx"`
	BoxOffice float64
}
