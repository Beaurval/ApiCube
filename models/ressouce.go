package models

import (
	"gorm.io/gorm"
)

//Ressource publiée par le citoyen
type Ressource struct {
	gorm.Model
	Titre             string
	Vues              int
	Votes             int
	Contenu           string
	TypeRessourceID   uint
	TypeRelationID    uint
	CitoyenID         uint
	ValidationAdmin   *bool
	Citoyen           Citoyen
	CategorieID       uint
	Categorie         Categorie
	Commentaires      []Commentaire
	Tags              []Tag `gorm:"many2many:tags_ressources;"`
	ActionsRessources []ActionRessource
	CitoyenVoted      []Citoyen `gorm:"many2many:ressources_voted;"`
}

//CreateRessourceInput model de création de ressource
type CreateRessourceInput struct {
	gorm.Model
	Titre           string `binding:"required"`
	Contenu         string `binding:"required"`
	CitoyenID       uint   `binding:"required"`
	TypeRelationID  uint   `binding:"required"`
	TypeRessourceID uint   `binding:"required"`
	CategorieID     uint   `binding:"required"`
}

//UpdateRessourceInput model pour mettre à jour la ressource
type UpdateRessourceInput struct {
	Titre           string
	Vues            int
	Votes           int
	ValidationAdmin *bool `gorm:"type:boolean"`
	Contenu         string
	TypeRelationID  uint
	CitoyenID       uint
	CategorieID     uint
	TypeRessourceID uint
}

// TableName sets the insert table name for this struct type
func (c *Ressource) TableName() string {
	return "ressources"
}
