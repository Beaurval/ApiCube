package main

import (
	"ApiCubes/controllers"
	"ApiCubes/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	//Routes citoyens
	r.GET("/citoyens", controllers.FindCitoyens)
	r.GET("/citoyens/:id", controllers.FindCitoyen)
	r.PATCH("/citoyens/:id", controllers.UpdateCitoyen)
	r.POST("/citoyens", controllers.CreateCitoyen)
	r.DELETE("/citoyens/:id", controllers.DeleteCitoyen)

	//Routes ressources
	r.GET("/ressources", controllers.FindRessources)
	r.GET("/ressources/:id", controllers.FindRessource)
	r.PATCH("/ressources/:id", controllers.UpdateRessource)
	r.DELETE("/ressources/:id", controllers.DeleteRessource)
	r.POST("/ressources", controllers.CreateRessource)
	r.Run()
}
