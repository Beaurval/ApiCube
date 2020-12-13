package models

import (
	"gorm.io/gorm"
)

//Tag renseigne les sujets des ressources
type Tag struct {
	gorm.Model
	Nom        string
	Ressources []Ressource `gorm:"many2many:tags_ressources;"`
}

//CreateTagInput model d'ajout de tag
type CreateTagInput struct {
	gorm.Model
	Nom string `json:"nom" binding:"required"`
}

//UpdateTagInput model d'update de tag
type UpdateTagInput struct {
	gorm.Model
	Nom string `json:"nom" binding:"required"`
}

// TableName sets the insert table name for this struct type
func (r *Tag) TableName() string {
	return "tags"
}
