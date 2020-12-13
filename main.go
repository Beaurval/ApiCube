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

	//Routes ressources
	r.GET("/commentaires", controllers.FindCommentaires)
	r.GET("/commentaires/:id", controllers.FindCommentaire)
	r.PATCH("/commentaires/:id", controllers.UpdateCommentaire)
	r.DELETE("/commentaires/:id", controllers.DeleteCommentaire)
	r.POST("/commentaires", controllers.CreateCommentaire)

	//Routes type de relation
	r.GET("/typeRelations", controllers.FindTypeRelations)
	r.GET("/typeRelations/:id", controllers.FindTypeRelation)
	r.PATCH("/typeRelations/:id", controllers.UpdateTypeRelation)
	r.DELETE("/typeRelations/:id", controllers.DeleteTypeRelation)
	r.POST("/typeRelations", controllers.CreateTypeRelation)

	//Routes rang
	r.GET("/rangs", controllers.FindRangs)
	r.GET("/rangs/:id", controllers.FindRang)
	r.PATCH("/rangs/:id", controllers.UpdateRang)
	r.DELETE("/rangs/:id", controllers.DeleteRang)
	r.POST("/rangs", controllers.CreateRang)
	r.Run()
}
