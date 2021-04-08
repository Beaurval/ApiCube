package main

import (
	"ApiCubes/controllers"
	middleware "ApiCubes/middleware"
	"ApiCubes/models"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/fichiers", http.Dir("fichiers"))

	models.ConnectDataBase()

	authMiddleware := middleware.InitAuth()
	//adminMiddleware := middleware.InitAdmin()

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	api := r.Group("/api")

	// Refresh time can be longer than token timeout
	api.GET("/refresh_token", authMiddleware.RefreshHandler)

	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
			AllowHeaders:     []string{"Origin", "content-type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		models.ConnectDataBase()

		//Routes citoyens

		api.GET("/citoyens/:id", controllers.FindCitoyen)
		api.PATCH("/citoyens/:id", controllers.UpdateCitoyen)
		api.POST("/citoyens", controllers.CreateCitoyen)
		api.DELETE("/citoyens/:id", controllers.DeleteCitoyen)

		//Routes fichier
		api.POST("/upload/:id", controllers.Upload)
		api.GET("/files", controllers.FindFiles)
		api.GET("/files/:id", controllers.FindFile)
		api.DELETE("/files/:id", controllers.DeleteFile)

		//Routes ressources
		api.GET("/ressources", controllers.FindRessources)
		api.GET("/ressources/:id", controllers.FindRessource)
		api.PATCH("/ressources/:id", controllers.UpdateRessource)
		api.DELETE("/ressources/:id", controllers.DeleteRessource)
		api.DELETE("/ressources/:id/tags/:idTag", controllers.DeleteTagRessource)
		api.DELETE("/ressources/:id/action/:idCitoyen", controllers.DeleteActionRessource)
		api.POST("/ressources", controllers.CreateRessource)
		api.POST("/ressources/tags/:id/:idTag", controllers.AddTagRessource)
		api.POST("/ressources/action/:id/:idCitoyen", controllers.AddActionRessource)

		//Routes commentaires
		api.GET("/commentaires", controllers.FindCommentaires)
		api.GET("/commentaires/:id", controllers.FindCommentaire)
		api.PATCH("/commentaires/:id", controllers.UpdateCommentaire)
		api.DELETE("/commentaires/:id", controllers.DeleteCommentaire)
		api.POST("/commentaires", controllers.CreateCommentaire)

		//Routes type de relation
		api.GET("/typeRelations", controllers.FindTypeRelations)
		api.GET("/typeRelations/:id", controllers.FindTypeRelation)
		api.PATCH("/typeRelations/:id", controllers.UpdateTypeRelation)
		api.DELETE("/typeRelations/:id", controllers.DeleteTypeRelation)
		api.POST("/typeRelations", controllers.CreateTypeRelation)

		//Routes rang
		api.GET("/rangs", controllers.FindRangs)
		api.GET("/rangs/:id", controllers.FindRang)
		api.PATCH("/rangs/:id", controllers.UpdateRang)
		api.DELETE("/rangs/:id", controllers.DeleteRang)
		api.POST("/rangs", controllers.CreateRang)

		//Routes rang
		api.GET("/categories", controllers.FindCategories)
		api.GET("/categories/:id", controllers.FindCategorie)
		api.POST("/categories", controllers.CreateCategorie)

		//Routes tag
		api.GET("/tags", controllers.FindTags)
		api.GET("/tags/:id", controllers.FindTag)
		api.PATCH("/tags/:id", controllers.UpdateTag)
		api.DELETE("/tags/:id", controllers.DeleteTag)
		api.POST("/tags", controllers.CreateTag)

		//Routes typeRessources
		api.GET("/typeRessources", controllers.FindTypeRessources)
		api.GET("/typeRessources/:id", controllers.FindTypeRessource)
		api.PATCH("/typeRessources/:id", controllers.UpdateTypeRessource)
		api.DELETE("/typeRessources/:id", controllers.DeleteTypeRessource)
		api.POST("/typeRessources", controllers.CreateTypeRessource)

		//Routes votes
		api.POST("/voteRessources", controllers.VoterRessource)
		api.POST("/voteCommentaire", controllers.VoterCommentaire)
		api.DELETE("/voteRessources/:idCitoyen/:idRessource", controllers.RetirerVoteRessource)
		api.DELETE("/voteCommentaire/:idCitoyen/:idCommentaire", controllers.RetirerVoteCommentaire)

		//Routes relations
		api.GET("/relations/:id", controllers.FindRelationsDuCitoyen)
		api.GET("/inrelations/:id", controllers.FindRelationsOuEstLeCitoyen)
		api.POST("/relations", controllers.AjouterRelation)
		api.DELETE("/relations/:id", controllers.DeleteRelation)
	}

	authMiddleware.Authorizator = middleware.Authorizator

	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/citoyens", controllers.FindCitoyens)
	}

	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}
