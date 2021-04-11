package models

import (
	"gorm.io/gorm"
)

//Citoyen Utilisateur de l'application
type Citoyen struct {
	gorm.Model
	Adresse           string
	CodePostal        string
	Genre             string
	Mail              string
	MotDePasse        string
	Nom               string
	Prenom            string
	Pseudo            string
	Telephone         string
	Ville             string
	RangID            uint
	Rang              Rang
	Ressource         []Ressource
	RessourcesVoted   []Ressource       `gorm:"many2many:ressources_voted;"`
	CommentairesVoted []Commentaire     `gorm:"many2many:commentaires_voted;"`
	RessourcesViewed  []Ressource       `gorm:"many2many:ressources_views;"`
	Relations         []RelationCitoyen `gorm:"foreignKey:CitoyenID"`
	InRelations       []RelationCitoyen `gorm:"foreignKey:CitoyenCibleID"`
}

//CreateCitoyenInput model de création de citoyen
type CreateCitoyenInput struct {
	Adresse    string `binding:"required"`
	CodePostal string `binding:"required"`
	Genre      string `binding:"required"`
	Mail       string `binding:"required"`
	MotDePasse string `binding:"required"`
	Nom        string `binding:"required"`
	Prenom     string `binding:"required"`
	Pseudo     string `binding:"required"`
	Telephone  string `binding:"required"`
	Ville      string `binding:"required"`
}

//UpdateCitoyenInput model de création de citoyen
type UpdateCitoyenInput struct {
	gorm.Model
	Adresse    string
	CodePostal string
	Genre      string
	Mail       string
	MotDePasse string
	Nom        string
	Prenom     string
	Pseudo     string
	Telephone  string
	Ville      string
	RangID     uint
}

// TableName sets the insert table name for this struct type
func (c *Citoyen) TableName() string {
	return "citoyens"
}
