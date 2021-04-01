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
	TypeRessource     TypeRessource
	TypeRelationID    uint
	Relation          TypeRelation `gorm:"foreignkey:ID"`
	CitoyenID         uint
	Commentaires      []Commentaire
	Tags              []Tag             `gorm:"many2many:tags_ressources;"`
	ActionsRessources []ActionRessource `gorm:"foreignkey:RessourceID"`
}

//CreateRessourceInput model de création de ressource
type CreateRessourceInput struct {
	gorm.Model
	Titre          string `json:"titre" binding:"required"`
	Contenu        string `json:"contenu" binding:"required"`
	CitoyenID      uint   `json:"citoyenId" binding:"required"`
	TypeRelationID uint   `json:"typeRelationId" binding:"required"`
}

//UpdateRessourceInput model pour mettre à jour la ressource
type UpdateRessourceInput struct {
	gorm.Model
	Titre          string `json:"titre"`
	Vues           int    `json:"vues"`
	Votes          int    `json:"votes"`
	Contenu        string `json:"contenu"`
	TypeRelationID uint   `json:"typeRelationId"`
	CitoyenID      uint   `json:"citoyenId"`
}

// TableName sets the insert table name for this struct type
func (c *Ressource) TableName() string {
	return "ressources"
}
