package models

import (
	"gorm.io/gorm"
)

//TypeRelation Type de relation
type TypeRelation struct {
	gorm.Model
	Nom       string
	Ressource []Ressource
}

//CreateTypeRelationInput model de création de type de relation
type CreateTypeRelationInput struct {
	gorm.Model
	Nom string `binding:"required"`
}

//UpdateTypeRelationInput model d'update type de relation
type UpdateTypeRelationInput struct {
	gorm.Model
	Nom string `json:"nom"`
}

// TableName sets the insert table name for this struct type
func (c *TypeRelation) TableName() string {
	return "typeRelations"
}
