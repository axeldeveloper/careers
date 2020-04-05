package models

import (
	"github.com/jinzhu/gorm"
)

//Model Hero
type Hero struct {
	gorm.Model
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	FullName     string `json:"full_name"`
	Intelligence string `json:"intelligence"`
	Power        string `json:"power"`
	Occupation   string `json:"occupation"`
	Image        string `json:"image"`
}
