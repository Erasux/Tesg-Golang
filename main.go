package main

import (
	"SamirGG/Tesg-Golang/database"
	"SamirGG/Tesg-Golang/handlers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	fmt.Println("Conectando a la base de datos...")
	if err := database.ConnectToDatabase(); err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	defer database.DisconnectFromDatabase()

	router := gin.Default()

	// Rutas de eventos

	router.POST("/events", handlers.CreateEvent)
	router.GET("/events", handlers.FindEvents)
	router.GET("/events/:id", handlers.FindEventById)
	router.PUT("/events/:id", handlers.UpdateEvent)
	router.DELETE("/events/:id", handlers.DeleteEvent)

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
	fmt.Println("Servidor iniciado en el puerto 8080")
}
