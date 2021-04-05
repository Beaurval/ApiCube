package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindCategories Get all categories
func FindCategories(c *gin.Context) {
	var categorie []models.Categorie
	models.DB.Find(&categorie)

	c.JSON(http.StatusOK, gin.H{"data": categorie})
}

//FindCategorie récupère la categorie correspondante à l'id passé en paramètre
func FindCategorie(c *gin.Context) {
	var categorie models.Categorie

	if err := models.DB.Where("id = ?", c.Param("id")).First(&categorie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categorie})
}

//CreateCategorie ajoute unee catégorie
func CreateCategorie(c *gin.Context) {
	//Validate input
	var input models.CreateCategorieInput

	fmt.Println(c.Keys)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create type relation
	rang := models.Categorie{
		Nom: input.Nom,
	}
	models.DB.Create(&rang)

	c.JSON(http.StatusOK, gin.H{"data": rang})
}
