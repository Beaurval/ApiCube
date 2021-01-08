package controllers

import (
	"ApiCubes/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//Upload Uploder un fichier
func Upload(c *gin.Context) {
	var fileBdd models.File
	var ressource models.Ressource

	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	println(file)
	if err != nil {
		println(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("fichiers/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename

	f, err := os.Stat("fichiers/" + filename)
	if err != nil {
		return
	}

	//Enregistrement en BDD
	fileBdd.Nom = filename
	fileBdd.Taille = int(f.Size())
	fileBdd.Path = models.ServerAddress + "fichiers/" + filename
	fileBdd.RessourceID = ressource.ID
	fileBdd.Ressource = ressource
	fileBdd.TypeFile = "f"

	models.DB.Create(&fileBdd)

	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

//FindFiles Récupérer tous les fichiers
func FindFiles(c *gin.Context) {
	var files []models.File
	models.DB.Preload("Ressource").Find(&files)

	c.JSON(http.StatusOK, gin.H{"data": files})
}

//FindFile récupère le fichier correspondant à l'id passé en paramètre
func FindFile(c *gin.Context) {
	var file models.File

	if err := models.DB.Preload("Ressource").Where("id = ?", c.Param("id")).First(&file).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": file})
}

// DeleteFile supprimer une ressource
func DeleteFile(c *gin.Context) {
	// Get model if exist
	var file models.File
	if err := models.DB.Where("id = ?", c.Param("id")).First(&file).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&file)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
