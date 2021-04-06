package controllers

import (
	"ApiCubes/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindCategories Get all categories
func FindCategories(c *gin.Context) {
	var categorie []models.Categorie
	models.DB.Find(&categorie)
	var result []map[string]interface{}

	for i := 0; i < len(categorie); i++ {
		var commentaires []models.Commentaire
		var ressources []models.Ressource

		models.DB.Raw("SELECT commentaires.id, commentaires.created_at, commentaires.updated_at,commentaires.deleted_at, "+
			"commentaires.parent_id, commentaires.citoyen_id, commentaires.ressource_id, commentaires.contenu, commentaires.vote "+
			"FROM beaurval_apiflutter.ressources "+
			"INNER JOIN commentaires on commentaires.ressource_id = ressources.id "+
			"WHERE ressources.categorie_id = @id", sql.Named("id", categorie[i].ID)).Scan(&commentaires)
		models.DB.Table("ressources").Where("categorie.id = ?", categorie[i].ID).Select("*").Joins("left join categorie on ressources.categorie_id = categorie.id").Order("ressources.created_at DESC").Scan(&ressources)

		var lastRessource models.Ressource
		if len(ressources) > 0 {
			lastRessource = ressources[0]
		}

		result = append(result, gin.H{
			"TotalCommentaires": len(commentaires),
			"TotalRessources":   len(ressources),
			"LastRessource":     lastRessource,
			"Categorie":         categorie[i]})
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

//FindCategorie récupère la categorie correspondante à l'id passé en paramètre
func FindCategorie(c *gin.Context) {
	var categorie models.Categorie
	var commentaires []models.Commentaire
	var ressources []models.Ressource

	if err := models.DB.Where("id = ?", c.Param("id")).First(&categorie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Raw("SELECT commentaires.id, commentaires.created_at, commentaires.updated_at,commentaires.deleted_at, "+
		"commentaires.parent_id, commentaires.citoyen_id, commentaires.ressource_id, commentaires.contenu, commentaires.vote "+
		"FROM beaurval_apiflutter.ressources "+
		"INNER JOIN commentaires on commentaires.ressource_id = ressources.id "+
		"WHERE ressources.categorie_id = @id", sql.Named("id", categorie.ID)).Scan(&commentaires)
	models.DB.Table("ressources").Where("categorie.id = ?", c.Param("id")).Select("*").Joins("left join categorie on ressources.categorie_id = categorie.id").Order("ressources.created_at DESC").Scan(&ressources)

	var lastRessource models.Ressource
	if len(ressources) > 0 {
		lastRessource = ressources[0]
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{
		"Categorie":         categorie,
		"TotalRessources":   len(ressources),
		"TotalCommentaires": len(commentaires),
		"LastRessource":     lastRessource,
	}})
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
		Nom:         input.Nom,
		Description: input.Description,
	}
	models.DB.Create(&rang)

	c.JSON(http.StatusOK, gin.H{"data": rang})
}
