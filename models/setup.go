package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB base de données
var DB *gorm.DB

const ServerAddress = "http://localhost:8080/"

//ConnectDataBase se connecte à la base de données renseignée dans la chaine de connexion
func ConnectDataBase() {
	dsn := "beaurval:alwayscesi@tcp(mysql-beaurval.alwaysdata.net)/beaurval_apiflutter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&File{}, &Tag{}, &TypeRessource{}, &Rang{}, &Citoyen{}, &TypeRelation{}, &Ressource{}, &Commentaire{}, &ActionRessource{})

	DB = db
}
