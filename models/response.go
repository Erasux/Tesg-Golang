package models

// ErrorResponse representa una respuesta de error de la API
type ErrorResponse struct {
	Error   string `json:"error" example:"Mensaje de error"`
	Details string `json:"details,omitempty" example:"Detalles adicionales del error"`
	Field   string `json:"field,omitempty" example:"Campo que causó el error"`
}

// SuccessResponse representa una respuesta exitosa de la API
type SuccessResponse struct {
	Message string `json:"message" example:"Operación realizada exitosamente"`
}

// CheckEventsResponse representa la respuesta del endpoint de verificación de eventos
type CheckEventsResponse struct {
	Message       string  `json:"message" example:"Eventos actualizados exitosamente"`
	Details       string  `json:"details" example:"Se actualizó el estado de gestión para eventos revisados"`
	UpdatedEvents []Event `json:"updatedEvents,omitempty"`
}
