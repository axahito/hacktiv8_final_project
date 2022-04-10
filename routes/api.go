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

	// photoRoutes := route.Group("/photo")
	// {
	// 	route.GET("/")
	// 	route.GET("/:photo")
	// 	route.POST("/")
	// 	route.PUT("/:photo")
	// 	route.DELETE("/:photo")
	// }

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
