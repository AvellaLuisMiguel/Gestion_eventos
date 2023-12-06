package models

import (
)

type Event struct {
	ID            int    `bson:"id_event"`
	Name          string `bson:"name_event"`
	Type          string `bson:"type_event"`
	Description   string `bson:"description_event"`
	Date          string `bson:"date_event"`
	State         int    `bson:"state_event"`
	Clasification string `bson:"clasification_event"`
}

func NewEvent() *Event{
	return &Event{}
}

func (e *Event) ToStringEvent()string{
	event:="{"+"Name: "+e.Name+", Type: "+e.Type+", Description"+e.Description+", Date: "+e.Date+"}"
	return event
}