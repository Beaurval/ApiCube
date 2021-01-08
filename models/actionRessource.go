package models

import "gorm.io/gorm"

//ActionRessource définit les actions de l'utilisateur sur la ressource (favoris ect...)
type ActionRessource struct {
	gorm.Model
	RessourceID int
	Ressource   Ressource
	CitoyenID   int
	Citoyen     Citoyen
	Favoris     bool
	MisDeCote   bool
	Exploite    bool
}

//CreateActionRessourceInput model de création d'action de ressource
type CreateActionRessourceInput struct {
	gorm.Model
	Favoris   bool `json:"favoris"`
	MisDeCote bool `json:"misDeCote"`
	Exploite  bool `json:"exploite"`
}
