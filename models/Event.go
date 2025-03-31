package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive" //Importación de MongoDB para el ID
)

// Tipo de estado de evento
type EventStatus string

const (
	StatusPending  EventStatus = "Pendiente por revisar"
	StatusReviewed EventStatus = "Revisado"
)

// Tipo de estado de gestión
type ManagementStatus string

const (
	ManagementRequired    ManagementStatus = "Requiere gestión"
	ManagementNotRequired ManagementStatus = "Sin gestión"
	ManagementUndefined   ManagementStatus = "" //Estado inicial por defecto
)

// Estructura de evento
type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB usa _id
	Name        string             `bson:"name" json:"name" binding:"required"`
	EventType   string             `bson:"eventType" json:"eventType" binding:"required"`
	Description string             `bson:"description" json:"description"`
	Date        time.Time          `bson:"date" json:"date" binding:"required"`
	Status      EventStatus        `bson:"status" json:"status"`

	ManagementStatus ManagementStatus `bson:"managementStatus,omitempty" json:"managementStatus,omitempty"`
}
