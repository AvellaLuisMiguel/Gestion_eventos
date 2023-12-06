package main

import (
	handlers "Events/api/handlers"
	"Events/api/routes"
	"Events/db"

	"Events/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dataBase := db.NewDatabase()
	logs:= docs.NewLogs("logs")
	logs.GenerateLogs()
	dataBase.Init()
	//dataBase.PrintCollections()
	eventHandler := handlers.NewEventHandler(*dataBase.GetModel(),*logs)
	routes.Setup(router, eventHandler)
	router.Run(":3000")
}
