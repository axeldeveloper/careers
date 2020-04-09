package models

import (
	u "levpay/utils"

	uuid "github.com/satori/go.uuid"
)

//Model Hero
type Heroe struct {
	//gorm.Model
	ID           uuid.UUID    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name         string       `sql:"type:VARCHAR(90);not null;unique"  json:"name"`
	FullName     string       `sql:"type:VARCHAR(150);not null"        json:"fullname"`
	Intelligence int          `sql:"type:int;not null"         		   json:"intelligence"`
	Power        int          `sql:"type:int;not null"                 json:"power"`
	Occupation   string       `sql:"type:VARCHAR(90);not null"         json:"occupation"`
	Image        string       `sql:"type:VARCHAR(750);not null"		   json:"image"`
	Relatives    uint32       `sql:"type:int " 			 			   json:"relatives"`
	Type         string       `sql:"type:super_type"    			   json:"type"`
	Connection   []Connection `json:"connections"` // one-to-many relationship
}

type Connection struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	HeroeID uuid.UUID `gorm:"type:uuid;            json:"heroe_id"`
	Group   string    `sql:"type:VARCHAR(90)"      json:"group"`
}

// set User's table name to be `Herois`
// func (User) TableName() string {
//	return "Herois"
// }

// Disable table name's pluralization globally
//db.SingularTable(true)

//Validate Model
func (hero *Heroe) Validate() (map[string]interface{}, bool) {

	var heroe Heroe

	err := GetDB().Debug().Model(&Heroe{}).Where("name = ?", hero.Name).Take(&heroe).Error

	if err != nil {
		return u.Message(false, "Error checked name"), false
	}

	if heroe.Name == hero.Name {
		return u.Message(false, "Name is duplicated"), false
	}

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
	var err error
	if resp, ok := hero.Validate(); !ok {
		return resp
	}

	err = GetDB().Create(hero).Error

	if err != nil {
		///return &Heroe{}, db.Error
		return u.Message(false, "Error create super heroes")
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

func (hero *Heroe) GetHeroes() (*[]Heroe, error) {

	heroes := []Heroe{}

	err := GetDB().Debug().Preload("Connection").Model(&Heroe{}).Limit(100).Find(&heroes).Error
	if err != nil {
		return &[]Heroe{}, err
	}

	return &heroes, nil
}

func (hero *Heroe) GetHeroe(pid string) (*Heroe, error) {
	err := GetDB().Debug().Preload("Connection").Model(&Heroe{}).Where("id = ?", pid).Take(&hero).Error

	if err != nil {
		return &Heroe{}, nil
	}
	return hero, nil
}

func (hero *Heroe) GetNameHeroe(name string) (*Heroe, error) {

	var c = Connection{}
	err := GetDB().Debug().Model(&Heroe{}).Related(&Connection{}).Where("name = ?", name).Take(&hero, &c).Error

	if err != nil {
		return &Heroe{}, nil
	}
	return hero, nil
}
