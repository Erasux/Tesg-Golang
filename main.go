package main

import (
	"SamirGG/Tesg-Golang/database"
	"SamirGG/Tesg-Golang/handlers"
	"SamirGG/Tesg-Golang/models"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Conectar a la base de datos
	fmt.Println("Conectando a la base de datos...")
	if err := database.ConnectToDatabase(); err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	defer database.DisconnectFromDatabase()

	// Verificar si hay eventos y si no, crear datos iniciales
	events, err := database.FindEvents()
	if err != nil {
		log.Printf("Error al verificar eventos: %v", err)
	} else if len(events) == 0 {
		fmt.Println("No hay eventos en la base de datos. Creando datos iniciales...")
		seedDataBase()
	}

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

func seedDataBase() {
	events := []models.Event{
		{
			ID:               primitive.NewObjectID(),
			Name:             "Evento 1",
			EventType:        "Tipo de evento 1",
			Description:      "Descripción del evento 1",
			Date:             time.Now(),
			Status:           models.StatusPending,
			ManagementStatus: models.ManagementUndefined,
		},
		{
			ID:               primitive.NewObjectID(),
			Name:             "Evento 2",
			EventType:        "Tipo de evento 2",
			Description:      "Descripción del evento 2",
			Date:             time.Now().AddDate(0, 0, 1),
			Status:           models.StatusPending,
			ManagementStatus: models.ManagementUndefined,
		},
		{
			ID:               primitive.NewObjectID(),
			Name:             "Evento 3",
			EventType:        "Tipo de evento 3",
			Description:      "Descripción del evento 3",
			Date:             time.Now().AddDate(0, 0, 2),
			Status:           models.StatusPending,
			ManagementStatus: models.ManagementUndefined,
		},
	}

	for _, event := range events {
		database.InsertEvent(event)
	}
}
