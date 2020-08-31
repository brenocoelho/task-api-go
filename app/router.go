package app

import (
	"task-api/controllers"
	"task-api/middlewares"

	"github.com/gin-gonic/gin"
)

func route() {

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(middlewares.CORSMiddleware())

	router.POST("/api/v1/login", controllers.Login)
	router.POST("/api/v1/signup", controllers.CreateUser)

	router.GET("/api/v1/tasks", middlewares.TokenAuthMiddleware(), controllers.FindTasks)
	router.POST("/api/v1/tasks", middlewares.TokenAuthMiddleware(), controllers.CreateTask)
	router.GET("/api/v1/tasks/:id", middlewares.TokenAuthMiddleware(), controllers.FindTask)
	router.PUT("/api/v1/tasks/:id", middlewares.TokenAuthMiddleware(), controllers.UpdateTask)
	router.DELETE("/api/v1/tasks/:id", middlewares.TokenAuthMiddleware(), controllers.DeleteTask)

	router.GET("/api/v1/tags", middlewares.TokenAuthMiddleware(), controllers.FindTags)
	router.POST("/api/v1/tags", middlewares.TokenAuthMiddleware(), controllers.CreateTag)
	router.GET("/api/v1/tags/:id", middlewares.TokenAuthMiddleware(), controllers.FindTag)
	router.PUT("/api/v1/tags/:id", middlewares.TokenAuthMiddleware(), controllers.UpdateTag)
	router.DELETE("/api/v1/tags/:id", middlewares.TokenAuthMiddleware(), controllers.DeleteTag)
}
