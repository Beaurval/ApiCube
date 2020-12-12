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

// TableName sets the insert table name for this struct type
func (c *TypeRelation) TableName() string {
	return "typeRelations"
}
