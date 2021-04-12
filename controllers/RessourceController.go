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
	var result []models.RessourceDisplay

	models.DB.Preload("Commentaires").Preload("TypeRelation").Preload("CitoyenViewedRessource").Preload("Categorie").Preload("CitoyenVoted").Preload("Citoyen").Preload("Tags").Find(&ressources)

	for i := 0; i < len(ressources); i++ {
		result = append(result, models.RessourceDisplay{
			ID:                ressources[i].ID,
			DeletedAt:         ressources[i].DeletedAt.Time,
			CreatedAt:         ressources[i].CreatedAt,
			UpdatedAt:         ressources[i].UpdatedAt,
			Titre:             ressources[i].Titre,
			Contenu:           ressources[i].Contenu,
			Vues:              len(ressources[i].CitoyenViewedRessource),
			Votes:             len(ressources[i].CitoyenVoted),
			CommentairesCount: len(ressources[i].Commentaires),
			TypeRessourceID:   ressources[i].TypeRessourceID,
			TypeRelationID:    ressources[i].TypeRelationID,
			TypeRelation:      ressources[i].TypeRelation,
			CitoyenID:         ressources[i].CitoyenID,
			ValidationAdmin:   ressources[i].ValidationAdmin,
			Citoyen:           ressources[i].Citoyen,
			CategorieID:       ressources[i].CategorieID,
			Categorie:         ressources[i].Categorie,
			Commentaires:      ressources[i].Commentaires,
			Tags:              ressources[i].Tags,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

//FindRessource récupère la ressource correspondant à l'id passé en paramètre
func FindRessource(c *gin.Context) {
	var ressource models.Ressource

	if err := models.DB.Preload("CitoyenViewedRessource").Preload("CitoyenVoted").Preload("Categorie").Preload("Commentaires.Citoyen").Preload("Commentaires.CitoyenVoted").Preload("Citoyen").Preload("Tags").Preload("Commentaires").Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var model = models.RessourceDisplay{
		ID:                ressource.ID,
		DeletedAt:         ressource.DeletedAt.Time,
		CreatedAt:         ressource.CreatedAt,
		UpdatedAt:         ressource.UpdatedAt,
		Titre:             ressource.Titre,
		Contenu:           ressource.Contenu,
		Vues:              len(ressource.CitoyenViewedRessource),
		Votes:             len(ressource.CitoyenVoted),
		CommentairesCount: len(ressource.Commentaires),
		TypeRessourceID:   ressource.TypeRessourceID,
		TypeRelationID:    ressource.TypeRelationID,
		CitoyenID:         ressource.CitoyenID,
		ValidationAdmin:   ressource.ValidationAdmin,
		Citoyen:           ressource.Citoyen,
		CategorieID:       ressource.CategorieID,
		Categorie:         ressource.Categorie,
		CitoyenVoted:      ressource.CitoyenVoted,
		Commentaires:      ressource.Commentaires,
		Tags:              ressource.Tags,
	}

	c.JSON(http.StatusOK, gin.H{"data": model})
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
