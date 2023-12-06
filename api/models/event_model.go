package models

import (

	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type EventModel struct {
	db *mongo.Collection
}

func NewEventModel(db *mongo.Database) *EventModel {
	return &EventModel{
		db: db.Collection("events"),
	}
}

func (r *EventModel) AddEvent(event *Event) error {
	_, err := r.db.InsertOne(context.Background(), event)
	return err
}

func (r *EventModel) GetAllEvents() ([]*Event, error) {
	filter := bson.D{}
	cursor, err := r.db.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	return r.listEvents(cursor)
}

func (r *EventModel) listEvents(cursor *mongo.Cursor) ([]*Event, error){
	var events []*Event
	for cursor.Next(context.Background()) {
		var event Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events,nil
}

func (r *EventModel) FindEventById(id_event int) (*Event, error) {
	filter := bson.D{{Key: "id_event", Value: id_event}}
	var result Event
	err := r.db.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No se enco|ntraron documentos con el ID especificado")
		}
		return nil, err
	}
	return &result, nil
}

func (r *EventModel) UpdateEvent(id_event int, updateEvents *Event) (*Event, error) {
	filter := bson.D{{Key: "id_event", Value: id_event}}
	update := bson.D{{Key: "$set", Value: updateEvents}}
	result := r.db.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No se encontraron documentos con el ID especificado")
		}
		return nil, result.Err()
	}
	var updatedDoc Event
	if err := result.Decode(&updatedDoc); err != nil {
		return nil, err
	}
	return &updatedDoc, nil
}

func (r *EventModel) DeleteEvent(id_event int) error {
	filter := bson.D{{Key: "id_event", Value: id_event}}
	result, err := r.db.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("No se encontraron documentos con el ID especificado")
	}
	return nil
}

func (r *EventModel) ManageEvent(id_event int, clasification string) error{
	filter := bson.D{{Key: "id_event", Value: id_event}}
	var update Event
	err := r.db.FindOne(context.Background(), filter).Decode(&update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("No se enco|ntraron documentos con el ID especificado")
		}
	}
	if update.State == 1 {
		return fmt.Errorf("No se puede modificar por el estado del evento")

	}
	return r.updateClasificationEvent(update,clasification,filter,err)
}

func (r *EventModel) updateClasificationEvent(update Event, clasification string, filter primitive.D, err error) error{
	update.Clasification = clasification
	fmt.Printf(update.Clasification)
	updatedEvent := update
	updateEvent := bson.D{{Key: "$set", Value: updatedEvent}}
	result := r.db.FindOneAndUpdate(context.Background(), filter, updateEvent, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			fmt.Errorf("No se encontraron documentos con el ID especificado")
		}
		result.Err()
	}
	return err
}

func (r *EventModel) PrintAllEvents(events[] *Event) string{
	var message string
	for _, event := range events {
		message+=event.ToStringEvent()+"\n"
	}
	return message
}