package main

import (
	"github.com/gin-gonic/gin"
	"github.com/s1ovac/simple-rest-api/internal/handler"
	"github.com/s1ovac/simple-rest-api/internal/storage"
)

func main() {
	router := gin.Default()
	memoryStorage := storage.NewMemoryStorage()
	handler := handler.NewHandler(memoryStorage)

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.GET("/employee", handler.GetAllEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()
}
