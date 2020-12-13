package models

import (
	"gorm.io/gorm"
)

//Rang définit le niveau de permission du citoyen
type Rang struct {
	gorm.Model
	Nom string
}

// TableName sets the insert table name for this struct type
func (r *Rang) TableName() string {
	return "rangs"
}
