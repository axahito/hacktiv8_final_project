package routes

import (
	"final_project/handlers"
	"final_project/middlewares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	route := gin.Default()
	store := cookie.NewStore([]byte("urgentMatters"))
	route.Use(sessions.Sessions("mysession", store))

	route.POST("/register", handlers.UserRegister)
	route.POST("/login", handlers.UserLogin)

	userRoutes := route.Group("/user")
	{
		userRoutes.Use(middlewares.Authenticate())
		userRoutes.PUT("/:user", handlers.UserUpdate)
		userRoutes.DELETE("/:user", handlers.UserDelete)
	}

	photoRoutes := route.Group("/photo")
	{
		photoRoutes.Use(middlewares.Authenticate())
		photoRoutes.GET("/", handlers.IndexPhoto)
		photoRoutes.GET("/:photo", handlers.ShowPhoto)
		photoRoutes.POST("/", handlers.CreatePhoto)
		photoRoutes.PUT("/:photo", handlers.PhotoUpdate)
		photoRoutes.DELETE("/:photo", handlers.PhotoDelete)
	}

	// commentRoute := route.Group("/comment")
	// {
	// 	route.GET("/")
	// 	route.GET("/:comment")
	// 	route.POST("/")
	// 	route.PUT("/:comment")
	// 	route.DELETE("/:comment")
	// }

	// socialRoutes := route.Group("/social")
	// {
	// 	route.GET("/")
	// 	route.GET("/:social")
	// 	route.POST("/")
	// 	route.PUT("/:social")
	// 	route.DELETE("/:social")
	// }

	return route
}
