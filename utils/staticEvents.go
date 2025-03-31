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
	// Crear fechas con formato específico (solo mes, día y año)
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	tomorrow := now.AddDate(0, 0, 1)
	dayAfterTomorrow := now.AddDate(0, 0, 2)

	events := []models.Event{
		{
			ID:               primitive.NewObjectID(),
			Name:             "Evento 1",
			EventType:        "Tipo de evento 1",
			Description:      "Descripción del evento 1",
			Date:             now,
			Status:           models.StatusPending,
			ManagementStatus: models.ManagementUndefined,
		},
		{
			ID:               primitive.NewObjectID(),
			Name:             "Evento 2",
			EventType:        "Tipo de evento 2",
			Description:      "Descripción del evento 2",
			Date:             tomorrow,
			Status:           models.StatusPending,
			ManagementStatus: models.ManagementUndefined,
		},
		{
			ID:               primitive.NewObjectID(),
			Name:             "Evento 3",
			EventType:        "Tipo de evento 3",
			Description:      "Descripción del evento 3",
			Date:             dayAfterTomorrow,
			Status:           models.StatusPending,
			ManagementStatus: models.ManagementUndefined,
		},
	}

	for _, event := range events {
		database.InsertEvent(event)
	}
}
