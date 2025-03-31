package handlers

import (
	"SamirGG/Tesg-Golang/database"
	"SamirGG/Tesg-Golang/models"
	"SamirGG/Tesg-Golang/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Crear un evento
// Input: evento
// Output: evento creado
func CreateEvent(c *gin.Context) {
	var event models.Event
	var err error

	// Obtener datos del formulario
	name := c.PostForm("name")
	eventType := c.PostForm("eventType")
	description := c.PostForm("description")
	dateStr := c.PostForm("date")

	// Validar campos requeridos
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El campo 'name' es requerido",
			"field": "name",
		})
		return
	}

	if eventType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El campo 'eventType' es requerido",
			"field": "eventType",
		})
		return
	}

	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El campo 'date' es requerido",
			"field": "date",
		})
		return
	}

	// Parsear la fecha
	event.Date, err = time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Formato de fecha inválido. Use el formato YYYY-MM-DD",
			"field": "date",
			"value": dateStr,
		})
		return
	}

	// Construir el evento
	event = models.Event{
		Name:             name,
		EventType:        eventType,
		Description:      description,
		Date:             time.Date(event.Date.Year(), event.Date.Month(), event.Date.Day(), 0, 0, 0, 0, time.UTC),
		Status:           models.StatusPending,
		ManagementStatus: models.ManagementUndefined,
	}

	// Insertar en la base de datos
	result, err := database.InsertEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al insertar el evento en la base de datos",
			"details": err.Error(),
		})
		return
	}

	event.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Evento creado exitosamente",
		"event":   event,
	})
}

// Buscar eventos
// Input: evento
// Output: evento encontrado
func FindEvents(c *gin.Context) {
	events, err := database.FindEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

// Buscar evento por ID
// Input: evento
// Output: evento encontrado
func FindEventById(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	event, err := database.FindEventById(objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// Actualizar un evento
// Input: evento
// Output: evento actualizado
func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = objectID
	result, err := database.UpdateEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Evento actualizado exitosamente"})
}

// Eliminar un evento
// Input: evento
// Output: evento eliminado
func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := database.DeleteEvent(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Evento eliminado exitosamente"})
}

// Revisar si hay eventos en la base de datos
func CheckEvents(c *gin.Context) {
	events, err := database.FindEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedEvents []models.Event
	for _, event := range events {
		updatedEvent := utils.ChangeManage(event)
		if updatedEvent.ManagementStatus != event.ManagementStatus {
			// Actualizar en la base de datos
			_, err := database.UpdateEvent(updatedEvent)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error al actualizar el evento",
					"eventId": event.ID,
					"details": err.Error(),
				})
				return
			}
			updatedEvents = append(updatedEvents, updatedEvent)
		}
	}

	if len(updatedEvents) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron eventos que necesiten actualización de gestión",
			"details": "Los eventos deben estar en estado 'Revisado' para tener un estado de gestión",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Eventos actualizados exitosamente",
		"details":       "Se actualizó el estado de gestión para eventos revisados",
		"updatedEvents": updatedEvents,
	})
}
