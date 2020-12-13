package models

import (
	"gorm.io/gorm"
)

//TypeRessource Type de ressource
type TypeRessource struct {
	gorm.Model
	Nom string
}

//CreateTypeRessourceInput model de création de type de ressource
type CreateTypeRessourceInput struct {
	gorm.Model
	Nom string `json:"nom"`
}

//UpdateTypeRessourceInput model d'update type de ressource
type UpdateTypeRessourceInput struct {
	gorm.Model
	Nom string `json:"nom"`
}

// TableName sets the insert table name for this struct type
func (c *TypeRessource) TableName() string {
	return "typeressources"
}
