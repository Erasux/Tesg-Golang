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

// Cambiar el estado de gestión de un evento
func ChangeManage(event models.Event) models.Event {
	if strings.ToLower(event.EventType) == "tipo de evento 1" || strings.ToLower(event.EventType) == "tipo de evento 2" {
		event.ManagementStatus = models.ManagementRequired
	} else {
		event.ManagementStatus = models.ManagementNotRequired
	}
	return event
}
