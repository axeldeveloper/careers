package models

import (
	"fmt"

	u "levpay/utils"

	"github.com/jinzhu/gorm"
)

//Model Hero
type Heroe struct {
	gorm.Model
	//ID           uint   `gorm:"primary_key"                      json:"id"`
	Name         string `sql:"type:VARCHAR(90);not null;unique"  json:"name"`
	FullName     string `sql:"type:VARCHAR(90);not null"         json:"fullname"`
	Intelligence string `sql:"type:VARCHAR(90);not null"         json:"intelligence"`
	Power        string `sql:"type:VARCHAR(90);not null"         json:"power"`
	Occupation   string `sql:"type:VARCHAR(90);not null"         json:"occupation"`
	Image        string `sql:"type:VARCHAR(750);not null"		 json:"image"`
}

// set User's table name to be `profiles`
//func (User) TableName() string {
//	return "profiles"
//}

// Disable table name's pluralization globally
//db.SingularTable(true)

//Validate Model
func (hero *Heroe) Validate() (map[string]interface{}, bool) {

	fmt.Print(hero)

	if hero.Name == "" {
		return u.Message(false, "Name is required"), false
	}
	if hero.FullName == "" {
		return u.Message(false, "Full Name is required"), false
	}
	if hero.Name == "" {
		return u.Message(false, "Name is required"), false
	}
	return u.Message(true, "success"), true
}

func (hero *Heroe) CreateHeroe() map[string]interface{} {

	if resp, ok := hero.Validate(); !ok {
		return resp
	}

	GetDB().Create(hero)

	if hero.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	resp := u.Message(true, "success")
	resp["hero"] = hero
	return resp
}

func (hero *Heroe) UpdateHeroe(uid uint64) (*Heroe, error) {
	db = GetDB().Debug().Model(&Heroe{}).Where("id = ?", uid).Take(&Heroe{}).UpdateColumns(
		map[string]interface{}{
			"Name":         hero.Name,
			"FullName":     hero.FullName,
			"Intelligence": hero.Intelligence,
			"Power":        hero.Power,
			"Occupation":   hero.Occupation,
			"Image":        hero.Image,
		},
	)
	if db.Error != nil {
		return &Heroe{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Heroe{}).Where("id = ?", uid).Take(&hero).Error
	if err != nil {
		return &Heroe{}, err
	}
	return hero, nil
}

func (hero *Heroe) DeleteHeroe(uid uint64) (int64, error) {

	//Soft Delete
	//db = GetDB().Debug().Model(&Heroe{}).Where("id = ?", uid).Take(&Heroe{}).Delete(&Heroe{})

	// Delete permanente
	db = GetDB().Debug().Unscoped().Delete(&Heroe{}, "id = ?", uid)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (hero *Heroe) GetHeroes() *[]Heroe {

	//var err error
	heroes := []Heroe{}

	err := GetDB().Debug().Model(&Heroe{}).Limit(100).Find(&heroes).Error
	if err != nil {
		return &[]Heroe{}
	}

	return &heroes
}

func (hero *Heroe) GetHeroe(pid uint64) *Heroe {
	//var err error
	err := GetDB().Debug().Model(&Heroe{}).Where("id = ?", pid).Take(&hero).Error

	if err != nil {
		return &Heroe{}
	}
	return hero
}
