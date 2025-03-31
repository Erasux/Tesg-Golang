package utils

import (
	"SamirGG/Tesg-Golang/models"
	"strings"
)

// Revisar si el evento está pendiente por revisar
func checkStatus(event models.Event) models.Event {
	if event.Status == models.StatusReviewed {
		return ChangeManage(event)
	}
	return event
}

// ChangeManage cambia el estado de gestión de un evento basado en su tipo
func ChangeManage(event models.Event) models.Event {
	// Solo establecer ManagementStatus si el evento está revisado
	if event.Status == models.StatusReviewed {
		eventType := strings.ToLower(event.EventType)
		if eventType == "tipo de evento 1" || eventType == "tipo de evento 2" {
			event.ManagementStatus = models.ManagementRequired
		} else {
			event.ManagementStatus = models.ManagementNotRequired
		}
	} else {
		// Si no está revisado, no debe tener ManagementStatus
		event.ManagementStatus = models.ManagementUndefined
	}
	return event
}
