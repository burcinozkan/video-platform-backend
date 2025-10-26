package main

import (
	_ "github.com/joho/godotenv/autoload"
	"video-platform-backend/internal/db"
	"video-platform-backend/internal/models"
	"video-platform-backend/internal/routes"
	"video-platform-backend/internal/storage"
)

func main() {
	db.ConnectDB()
	db.DB.AutoMigrate(&models.Task{})
	storage.InitS3()

	r := routes.SetupRouter()
	r.Run(":8080")
}
