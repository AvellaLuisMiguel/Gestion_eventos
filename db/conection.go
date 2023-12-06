package db

import (
	"Events/api/models"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct{
	model *models.EventModel
	client *mongo.Client
}

func NewDatabase() *Database{
	return &Database{}
}

func (d *Database) Init() {
	mongoURI := "mongodb://admin:12345@localhost:27017/events"
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	d.client = client
	db := client.Database("events")
	d.model = models.NewEventModel(db)
	log.Println("Conexión a MongoDB establecida correctamente")

}

func (d *Database) GetModel() *models.EventModel {
	return d.model
}

