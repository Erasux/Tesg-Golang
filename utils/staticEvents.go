package utils

import (
	"SamirGG/Tesg-Golang/database"
	"SamirGG/Tesg-Golang/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CheckEvents revisa si hay eventos en la base de datos
func CheckEvents() bool {
	events, err := database.FindEvents()
	if err != nil {
		log.Printf("Error al verificar eventos: %v", err)
	}
	return len(events) == 0
}

// SeedDataBase crea eventos estáticos en la base de datos
func SeedDataBase() {
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
