package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//FindRessources Récupérer toutes les ressources
func FindRessources(c *gin.Context) {
	var ressources []models.Ressource
	models.DB.Preload("Commentaires").Preload("Categorie").Preload("CitoyenVoted").Preload("Citoyen").Preload("Tags").Find(&ressources)

	c.JSON(http.StatusOK, gin.H{"data": ressources})
}

//FindRessource récupère la ressource correspondant à l'id passé en paramètre
func FindRessource(c *gin.Context) {
	var ressource models.Ressource

	if err := models.DB.Preload("CitoyenVoted").Preload("Categorie").Preload("Commentaires.Citoyen").Preload("Commentaires.CitoyenVoted").Preload("Citoyen").Preload("Tags").Preload("Commentaires").Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

//CreateRessource Ajoute une nouvelle ressource
func CreateRessource(c *gin.Context) {
	//Validate input
	var input models.CreateRessourceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Ressource
	ressource := models.Ressource{
		Titre:           input.Titre,
		Contenu:         input.Contenu,
		CitoyenID:       input.CitoyenID,
		TypeRelationID:  input.TypeRelationID,
		TypeRessourceID: input.TypeRessourceID,
		Vues:            0,
		Votes:           0,
		CategorieID:     input.CategorieID,
	}
	models.DB.Create(&ressource)

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

//AddTagRessource ajoute un tag à la ressource
func AddTagRessource(c *gin.Context) {
	// Get ressource if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Get tag if exist
	var tag models.Tag
	if err := models.DB.Where("id = ?", c.Param("idTag")).First(&tag).Error; err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&ressource).Association("Tags").Append([]models.Tag{tag})

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

//DeleteTagRessource supprime un tag de la ressource
func DeleteTagRessource(c *gin.Context) {
	// Get ressource if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Get tag if exist
	var tag models.Tag
	if err := models.DB.Where("id = ?", c.Param("idTag")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&ressource).Association("Tags").Delete([]models.Tag{tag})

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

//AddActionRessource ajoute un tag à la ressource
func AddActionRessource(c *gin.Context) {
	// Get ressource if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.CreateActionRessourceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idRessource, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// handle error
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	idCitoyen, err := strconv.Atoi(c.Param("idCitoyen"))
	if err != nil {
		// handle error
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&ressource).Association("ActionsRessources").Append(&models.ActionRessource{
		RessourceID: idRessource,
		CitoyenID:   idCitoyen,
		Favoris:     input.Favoris,
		MisDeCote:   input.MisDeCote,
		Exploite:    input.MisDeCote,
	})

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

//DeleteActionRessource supprime une action de ressource
func DeleteActionRessource(c *gin.Context) {
	// Get ressource if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Get actionRessource if exist
	var actionRessource models.ActionRessource
	if err := models.DB.Where("citoyen_id = ?", c.Param("idCitoyen")).Where("ressource_id = ?", c.Param("id")).First(&actionRessource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&ressource).Association("ActionsRessources").Delete([]models.ActionRessource{actionRessource})

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

// UpdateRessource update les informations d'une ressource
func UpdateRessource(c *gin.Context) {
	// Get model if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateRessourceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%v+\n", input)
	models.DB.Model(&ressource).Updates(map[string]interface{}{
		"Titre":           input.Titre,
		"Vues":            input.Vues,
		"Votes":           input.Votes,
		"ValidationAdmin": input.ValidationAdmin,
		"Contenu":         input.Contenu,
		"CitoyenID":       input.CitoyenID,
		"CategorieID":     input.CategorieID,
		"TypeRelationID":  input.TypeRelationID,
		"TypeRessourceID": input.TypeRessourceID,
	})

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

// DeleteRessource supprimer une ressource
func DeleteRessource(c *gin.Context) {
	// Get model if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&ressource)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
