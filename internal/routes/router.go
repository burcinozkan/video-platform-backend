package routes

import (
	"github.com/gin-gonic/gin"
	"video-platform-backend/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Task API
	api := r.Group("/tasks")
	{
		api.POST("/", handlers.CreateTask)
		api.GET("/", handlers.GetTasks)
		api.GET("/:id", handlers.GetTaskByID)
		api.PUT("/:id", handlers.UpdateTask)
		api.DELETE("/:id", handlers.DeleteTask)
	}

	r.POST("/videos/upload", handlers.UploadVideo)

	return r
}
