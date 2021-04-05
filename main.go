package main

import (
	"ApiCubes/controllers"
	"ApiCubes/models"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.StaticFS("/fichiers", http.Dir("fichiers"))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:44369"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	models.ConnectDataBase()

	//Routes citoyens
	r.GET("/citoyens", controllers.FindCitoyens)
	r.GET("/citoyens/:id", controllers.FindCitoyen)
	r.PATCH("/citoyens/:id", controllers.UpdateCitoyen)
	r.POST("/citoyens", controllers.CreateCitoyen)
	r.DELETE("/citoyens/:id", controllers.DeleteCitoyen)

	//Routes fichier
	r.POST("/upload/:id", controllers.Upload)
	r.GET("/files", controllers.FindFiles)
	r.GET("/files/:id", controllers.FindFile)
	r.DELETE("/files/:id", controllers.DeleteFile)

	//Routes ressources
	r.GET("/ressources", controllers.FindRessources)
	r.GET("/ressources/:id", controllers.FindRessource)
	r.PATCH("/ressources/:id", controllers.UpdateRessource)
	r.DELETE("/ressources/:id", controllers.DeleteRessource)
	r.DELETE("/ressources/:id/tags/:idTag", controllers.DeleteTagRessource)
	r.DELETE("/ressources/:id/action/:idCitoyen", controllers.DeleteActionRessource)
	r.POST("/ressources", controllers.CreateRessource)
	r.POST("/ressources/tags/:id/:idTag", controllers.AddTagRessource)
	r.POST("/ressources/action/:id/:idCitoyen", controllers.AddActionRessource)

	//Routes commentaires
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

	//Routes tag
	r.GET("/tags", controllers.FindTags)
	r.GET("/tags/:id", controllers.FindTag)
	r.PATCH("/tags/:id", controllers.UpdateTag)
	r.DELETE("/tags/:id", controllers.DeleteTag)
	r.POST("/tags", controllers.CreateTag)

	//Routes typeRessources
	r.GET("/typeRessources", controllers.FindTypeRessources)
	r.GET("/typeRessources/:id", controllers.FindTypeRessource)
	r.PATCH("/typeRessources/:id", controllers.UpdateTypeRessource)
	r.DELETE("/typeRessources/:id", controllers.DeleteTypeRessource)
	r.POST("/typeRessources", controllers.CreateTypeRessource)

	//Routes votes
	r.POST("/voteRessources", controllers.VoterRessource)
	r.POST("/voteCommentaire", controllers.VoterCommentaire)
	r.DELETE("/voteRessources/:idCitoyen/:idRessource", controllers.RetirerVoteRessource)
	r.DELETE("/voteCommentaire/:idCitoyen/:idCommentaire", controllers.RetirerVoteCommentaire)

	r.Run(":8081")
}
