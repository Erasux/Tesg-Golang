package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventStatus string // Define un tipo para el estado

const (
	StatusPending  EventStatus = "Pendiente por revisar"
	StatusReviewed EventStatus = "Revisado"
)

// Opcional: Para la clasificación
type ManagementStatus string

const (
	ManagementRequired    ManagementStatus = "Requiere gestión"
	ManagementNotRequired ManagementStatus = "Sin gestión"
	ManagementUndefined   ManagementStatus = "" // Estado inicial o si no está revisado
)

type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB usa _id
	Name        string             `bson:"name" json:"name" binding:"required"`
	EventType   string             `bson:"eventType" json:"eventType" binding:"required"`
	Description string             `bson:"description" json:"description"`
	Date        time.Time          `bson:"date" json:"date" binding:"required"`
	Status      EventStatus        `bson:"status" json:"status"`
	// Campo para la clasificación (Paso 6)
	ManagementStatus ManagementStatus `bson:"managementStatus,omitempty" json:"managementStatus,omitempty"`
}
