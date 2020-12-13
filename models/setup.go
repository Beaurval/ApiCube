package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB base de données
var DB *gorm.DB

//ConnectDataBase se connecte à la base de données renseignée dans la chaine de connexion
func ConnectDataBase() {
	dsn := "root:toor@tcp(127.0.0.1:3306)/projet_cube?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Tag{}, &Rang{}, &Citoyen{}, &TypeRelation{}, &Ressource{}, &Commentaire{})

	DB = db
}
