package models

import (
	"gorm.io/gorm"
)

//Categorie Utilisateur de l'application
type Categorie struct {
	gorm.Model
	Nom         string
	Description string
	Ressources  []Ressource
}

//CreateCategorieInput model de création de citoyen
type CreateCategorieInput struct {
	Nom         string
	Description string
}

// TableName sets the insert table name for this struct type
func (c *Categorie) TableName() string {
	return "categorie"
}
