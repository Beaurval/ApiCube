package models

import (
	"gorm.io/gorm"
)

//File Fichier associé à une ressource
type File struct {
	gorm.Model
	Nom         string
	Taille      int
	Path        string
	TypeFile    string
	RessourceID uint
	Ressource   Ressource
}
