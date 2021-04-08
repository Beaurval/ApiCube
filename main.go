package main

import (
	"ApiCubes/controllers"
	"ApiCubes/models"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type login struct {
	Mail     string `form:"Mail" json:"Mail" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
}

var identityKey = "Mail"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/fichiers", http.Dir("fichiers"))

	models.ConnectDataBase()

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 10,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var citoyen models.Citoyen

			if err := models.DB.Where("mail = ?", loginVals.Mail).First(&citoyen).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			userID := loginVals.Mail
			password := loginVals.Password

			if userID == citoyen.Mail && password == citoyen.MotDePasse {
				return &User{
					UserName:  userID,
					LastName:  citoyen.Nom,
					FirstName: citoyen.Prenom,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

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
		api.GET("/citoyens", controllers.FindCitoyens)
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

	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}
