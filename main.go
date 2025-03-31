package main

import (
	"SamirGG/Tesg-Golang/database"
	_ "SamirGG/Tesg-Golang/docs"
	"SamirGG/Tesg-Golang/handlers"
	"SamirGG/Tesg-Golang/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Eventos
// @version 1.0
// @description API para la gestión de eventos
// @host localhost:8080
// @BasePath /events
// @schemes http
// @consumes json
// @produces json
// @contact.name Samir Gonzallez
// @contact.email samirgg2000@gmail.com

// @tag.name Eventos
// @tag.description Operaciones relacionadas con eventos

// @definitions.ErrorResponse
// @definitions.SuccessResponse
// @definitions.CheckEventsResponse

func main() {
	//Conectar a la base de datos
	fmt.Println("Conectando a la base de datos...")
	if err := database.ConnectToDatabase(); err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	defer database.DisconnectFromDatabase()

	//Revisar que no haya eventos en la base de datos
	if utils.CheckEvents() {
		fmt.Println("No hay eventos en la base de datos. Creando datos iniciales...")
		utils.SeedDataBase()
	}

	router := gin.Default()

	// Rutas de eventos
	router.POST("/events", handlers.CreateEvent)
	router.GET("/events", handlers.FindEvents)
	router.GET("/events/:id", handlers.FindEventById)
	router.PUT("/events/:id", handlers.UpdateEvent)
	router.DELETE("/events/:id", handlers.DeleteEvent)
	router.GET("/events/check", handlers.CheckEvents)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
	fmt.Println("Servidor iniciado en el puerto 8080")
}
