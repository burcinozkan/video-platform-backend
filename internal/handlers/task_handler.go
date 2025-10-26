package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"video-platform-backend/internal/cache"
	"video-platform-backend/internal/db"
	"video-platform-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	status := c.Query("status")

	cacheKey := "tasks_all"
	if status != "" {
		cacheKey = "tasks_status_" + status
	}

	if cached, err := cache.Client.Get(cache.Ctx, cacheKey).Result(); err == nil {
		c.Data(http.StatusOK, "application/json", []byte(cached))
		return
	}

	if status != "" {
		db.DB.Where("status = ?", status).Find(&tasks)
	} else {
		db.DB.Find(&tasks)
	}

	jsonData, _ := json.Marshal(tasks)
	cache.Client.Set(cache.Ctx, cacheKey, jsonData, time.Minute*2)

	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := db.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := db.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updated models.Task
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&task).Updates(updated)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := db.DB.Delete(&models.Task{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
