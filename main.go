package main

import (
	"ApiCubes/controllers"
	"ApiCubes/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/citoyens", controllers.FindCitoyens)
	r.GET("/citoyens/:id", controllers.FindCitoyen)
	r.PATCH("/citoyens/:id", controllers.UpdateCitoyen)
	r.POST("/citoyens", controllers.CreateCitoyen)
	r.DELETE("/citoyens/:id", controllers.DeleteCitoyen)
	r.Run()
}
