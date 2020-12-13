package models

import (
	"gorm.io/gorm"
)

//Citoyen Utilisateur de l'application
type Citoyen struct {
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
	Rang       Rang
}

//CreateCitoyenInput model de création de citoyen
type CreateCitoyenInput struct {
	Adresse    string `json:"adresse" binding:"required"`
	CodePostal string `json:"cp" binding:"required"`
	Genre      string `json:"genre" binding:"required"`
	Mail       string `json:"mail" binding:"required"`
	MotDePasse string `json:"motDePasse" binding:"required"`
	Nom        string `json:"nom" binding:"required"`
	Prenom     string `json:"prenom" binding:"required"`
	Pseudo     string `json:"pseudo" binding:"required"`
	Telephone  string `json:"telephone" binding:"required"`
	Ville      string `json:"ville" binding:"required"`
	RangID     uint   `json:"RangID" binding:"required"`
}

//UpdateCitoyenInput model de création de citoyen
type UpdateCitoyenInput struct {
	gorm.Model
	Adresse    string `json:"adresse"`
	CodePostal string `json:"cp"`
	Genre      string `json:"genre"`
	Mail       string `json:"mail"`
	MotDePasse string `json:"motDePasse"`
	Nom        string `json:"nom"`
	Prenom     string `json:"prenom"`
	Pseudo     string `json:"pseudo"`
	Telephone  string `json:"telephone"`
	Ville      string `json:"ville"`
	RangID     uint   `json:"RangID"`
}

// TableName sets the insert table name for this struct type
func (c *Citoyen) TableName() string {
	return "citoyens"
}
