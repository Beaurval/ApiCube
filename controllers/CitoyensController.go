package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindCitoyens Get all citoyens
func FindCitoyens(c *gin.Context) {
	var citoyens []models.Citoyen
	models.DB.Preload("Rang").Find(&citoyens)

	c.JSON(http.StatusOK, gin.H{"data": citoyens})
}

//FindCitoyen récupère le citoyen correspondant à l'id passé en paramètre
func FindCitoyen(c *gin.Context) {
	var citoyen models.Citoyen

	if err := models.DB.Preload("Rang").Where("id = ?", c.Param("id")).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": citoyen})
}

// UpdateCitoyen update les informations d'un citoyen
func UpdateCitoyen(c *gin.Context) {
	// Get model if exist
	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", c.Param("id")).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateCitoyenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&citoyen).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": citoyen})
}

//CreateCitoyen Add a citoyen
func CreateCitoyen(c *gin.Context) {
	//Validate input
	var input models.CreateCitoyenInput

	fmt.Println(c.Keys)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create citoyen
	citoyen := models.Citoyen{
		Mail:       input.Mail,
		Ville:      input.Ville,
		CodePostal: input.CodePostal,
		Nom:        input.Nom,
		Prenom:     input.Prenom,
		MotDePasse: input.MotDePasse,
		Genre:      input.Genre,
		Adresse:    input.Adresse,
		Pseudo:     input.Pseudo,
		Telephone:  input.Telephone,
		RangID:     input.RangID,
	}
	models.DB.Create(&citoyen)

	c.JSON(http.StatusOK, gin.H{"data": citoyen})
}

// DeleteCitoyen supprimer un citoyen
func DeleteCitoyen(c *gin.Context) {
	// Get model if exist
	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", c.Param("id")).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&citoyen)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
