package models

import (
	u "levpay/utils"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Super represents the model for an super
type Super struct {
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
	SuperID uuid.UUID `gorm:"type:uuid;            json:"super_id"`
	Group   string    `sql:"type:VARCHAR(90)"      json:"group"`
}

// APIError example
type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

// @todo Validate Model
func (super *Super) Validate() (map[string]interface{}, bool) {

	checked := GetDB().Debug().Model(&Super{}).Where("name = ?", super.Name).First(&super).RecordNotFound()
	if !checked {
		return u.Message(false, "Name is duplicated"), false
	}

	//if heroe.Name == hero.Name {
	//	return u.Message(false, "Name is duplicated"), false
	//}

	if super.Name == "" {
		return u.Message(false, "Name is required"), false
	}
	if super.FullName == "" {
		return u.Message(false, "Full Name is required"), false
	}
	if super.Intelligence <= 0 {
		return u.Message(false, "Intelligence is required"), false
	}
	if super.Occupation == "" {
		return u.Message(false, "Occupation is required"), false
	}
	return u.Message(true, "success"), true
}

// CreateSuper godoc
// @Summary Create a new super
// @Description Create a new super with the input paylod
// @Tags supers
// @Accept  json
// @Produce  json
// @Param Super body Super true  "this is a test super"
// @Success 201 {object} Super
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Unable to register"
// @Router /api/supers [post]
func (super *Super) CreateSuper() (*Super, error) {
	err := GetDB().Create(super).Error
	if err != nil {
		return &Super{}, err
	}
	return super, nil
	// @Success 201 {object} string	"061bf710-9a64-4e36-8f04-d86806cb8cec"
}

// UpdateSuper godoc
// @Summary Update a super
// @Description Update a new super with the input paylod
// @Tags supers
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the super to be update"
// @Param Super body Super true  "this is a test Super"
// @Success 200 {object} Super
// @Success 202 {object} Super
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Unable to register"
// @Router /api/supers/{id} [put]
func (super *Super) UpdateSuper(uid string) (*Super, error) {
	db = GetDB().Debug().Model(&Super{}).Where("id = ?", uid).Take(&Super{}).UpdateColumns(
		map[string]interface{}{
			"Name":         super.Name,
			"FullName":     super.FullName,
			"Intelligence": super.Intelligence,
			"Power":        super.Power,
			"Occupation":   super.Occupation,
			"Image":        super.Image,
			"Relatives":    super.Relatives,
			"Type":         super.Type,
		},
	)
	if db.Error != nil {
		return &Super{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Preload("Connection").Model(&Super{}).Where("id = ?", uid).Take(&super).Error
	if err != nil {
		return &Super{}, err
	}
	return super, nil
}

// DeleteSuper godoc
// @Summary Delete super identified by the given id
// @Description Delete the super corresponding to the input id
// @Tags supers
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the super to be deleted"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Unable to register"
// @Router /api/supers/{id} [delete]
func (super *Super) DeleteSuper(uid string) (int64, error) {

	//Soft Delete
	//db = GetDB().Debug().Model(&Heroe{}).Where("id = ?", uid).Take(&Heroe{}).Delete(&Heroe{})

	// Delete permanente
	db = GetDB().Debug().Unscoped().Delete(&Super{}, "id = ?", uid)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// GetSupers godoc
// @Summary Get all supers
// @Description  return all supers
// @Tags supers
// @Accept  json
// @Produce  json
// @Success 200 {object} Super
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Unable to register"
// @Router /api/supers [get]
func (super *Super) GetSupers() (*[]Super, error) {
	supers := []Super{}
	err := GetDB().Debug().Preload("Connection").Model(&Super{}).Limit(100).Find(&supers).Error
	if err != nil {
		return &[]Super{}, err
	}
	//CloseDb()
	return &supers, nil
}

// GetSuper godoc
// @Summary Get details for a given super id
// @Description Get details of super corresponding to the input id
// @Tags supers
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the super"
// @Success 200 {object} Super
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Unable to register"
// @Router /api/supers/{id} [get]
func (super *Super) GetSuper(pid string) (*Super, error) {
	err := GetDB().Debug().Preload("Connection").Model(&Super{}).Where("id = ?", pid).Take(&super).Error
	if err != nil {
		return &Super{}, nil
	}
	//CloseDb()
	return super, nil
}

// GetSupersType godoc
// @Summary Get all supers [heroes and villains]
// @Description  return all supers heroes and villains
// @Tags HeroesOrVillains
// @Accept  json
// @Produce  json
// @Success 200 {object} Super
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Unable to register"
// @Param type path string true "type Hero or villain"
// @Router /api/supers/type/{type} [get]
func (super *Super) GetSupersType(t string) (*[]Super, error) {

	supers := []Super{}
	err := GetDB().Debug().Preload("Connection").Model(&Super{}).Where("type = ?", t).Limit(100).Find(&supers).Error

	if err != nil {
		return &[]Super{}, err
	}

	return &supers, nil
}

// GetNameSuper godoc
// @Summary Get details for a given super Name
// @Description Get details of super corresponding to the input name
// @Tags supers
// @Accept  json
// @Produce  json
// @Param name path string true "name of the super"
// @Success 200 {object} Super
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Unable to register"
// @Router /api/supers/name/{name} [get]
func (super *Super) GetNameSuper(name string) (*Super, error) {

	err := GetDB().Debug().Preload("Connection").Model(&Super{}).Where("name = ?", name).Take(&super).Error

	if err != nil {
		return &Super{}, nil
	}
	return super, nil
}
