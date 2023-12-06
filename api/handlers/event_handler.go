// Package handlers provides HTTP request handlers.
//
// This file contains the handlers for managing events.
//
//	Schemes: http
//	Host: localhost:8080
//	Version: 1.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- o
package handlers

import (
	// Import packege models str conv and gin
	"Events/api/models"
	"Events/docs"
	"strconv"
	"github.com/gin-gonic/gin"
)

// EventHandler manages events.
type EventHandler struct {
	event  models.EventModel
	writer docs.Logs
}

// NewEventHandler creates a new EventHandler.
func NewEventHandler(event models.EventModel, writer docs.Logs) *EventHandler {
	return &EventHandler{
		event:  event,
		writer: writer,
	}
}

// AddEvent adds a new event.
//
// swagger:operation POST /events addEvent
//
// Adds a new event to the system.
//
// Responses:
//
//	201: eventResponse
//	400: errorResponse
//	500: errorResponse
func (h *EventHandler) AddEvent(c *gin.Context) {
	var newEvent models.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		h.writer.WriteLog("error" + err.Error())
		return
	}
	if err := h.event.AddEvent(&newEvent); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		h.writer.WriteLog("error" + err.Error())
		return
	}
	c.JSON(201, gin.H{"message": "Evento creado exitosamente"})
	h.writer.WriteLog("Evento creado=> " + newEvent.ToStringEvent())
}

// UpdateEvent updates an existing event.
//
// swagger:operation PUT /events/{id} updateEvent
//
// Updates an existing event in the system.
//
// Responses:
//
//	200: eventResponse
//	400: errorResponse
//	500: errorResponse
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	var eventUpdate *models.Event
	id_event, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&eventUpdate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.event.UpdateEvent(id_event, eventUpdate)
	h.writer.WriteLog("Evento Actualizado=> "+eventUpdate.ToStringEvent())
}

// GetEvents gets a list of all events.
//
// swagger:operation GET /events getEvents
//
// Gets a list of all events in the system.
//
// Responses:
//
//	200: []eventResponse
//	500: errorResponse
func (h *EventHandler) GetEvents(c *gin.Context) {
	events, err := h.event.GetAllEvents()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		h.writer.WriteLog("error: " + err.Error())
		return
	}
	c.JSON(200, events)
	h.writer.WriteLog("Eventos=> "+h.event.PrintAllEvents(events))
}

// ManageEvent manages the classification of an event.
//
// swagger:operation DELETE /events/{id} manageEvent
//
// Manages the classification of an event in the system.
//
// Responses:
//
//	200: successResponse
//	400: errorResponse
//	500: errorResponse
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	id_event, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.event.DeleteEvent(id_event)
	h.writer.WriteLog("Se elimino el evento con id=>"+c.Param("id"))
}

// Get event by id.
//
// swagger:operation GET /events/{id} manageEvent
//
// Manages the classification of an event in the system.
//
// Responses:
//
//	200: successResponse
//	400: errorResponse
//	500: errorResponse
func (h *EventHandler) GetEventById(c *gin.Context) {
	id_event, err := strconv.Atoi(c.Param("id"))
	event, err := h.event.FindEventById(id_event)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, event)
	h.writer.WriteLog("Se encontró el evento con ID=> "+ event.ToStringEvent())
}

// ManageEvent manages the classification of an event.
//
// swagger:operation PUT /events/{id}/{classification} manageEvent
//
// Manages the classification of an event in the system.
//
// Responses:
//
//	200: successResponse
//	400: errorResponse
//	500: errorResponse
func (h *EventHandler) ManageEvent(c *gin.Context) {
	id_event, err := strconv.Atoi(c.Param("id"))
	clasification := c.Param("clasification")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = h.event.ManageEvent(id_event, clasification)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		h.writer.WriteLog("El evento no puede cambiar por su estado")
		return
	}
	c.JSON(201, gin.H{"message": "El evento cambio su clasificación"})
	h.writer.WriteLog("El evento cambió su clasificacion=> "+c.Param("id"))
}
