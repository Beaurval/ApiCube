package models

import (
	"gorm.io/gorm"
)

//Commentaire de ressource
type Commentaire struct {
	gorm.Model
	ParentID    *uint
	Reponses    []Commentaire `gorm:"foreignkey:ParentID"`
	RessourceID uint
	Contenu     string
	Vote        int
}

//CreateCommentaireInput model de création de ressource
type CreateCommentaireInput struct {
	gorm.Model
	ParentID    *uint  `json:"idParent"`
	RessourceID uint   `json:"idRessource" binding:"required"`
	Contenu     string `json:"contenu" binding:"required"`
}

//UpdateCommentaireInput model de mise à jour d'un commentaire de ressource
type UpdateCommentaireInput struct {
	gorm.Model
	Contenu string `json:"contenu"`
	Vote    int    `json:"vote"`
}

// TableName sets the insert table name for this struct type
func (c *Commentaire) TableName() string {
	return "commentaires"
}
