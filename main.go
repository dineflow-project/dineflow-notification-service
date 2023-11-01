package main

import (
	"fmt"
	"os"

	"github.com/dineflow-project/dineflow-notification-service/database"
	"github.com/dineflow-project/dineflow-notification-service/rabbitmq"
	"github.com/dineflow-project/dineflow-notification-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	router := gin.Default()

	err := database.ConnectToDB(os.Getenv("MONGO_URI"))
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	} else {
		fmt.Println("Connected to database")
	}

	go rabbitmq.StartNotificationConsumer(os.Getenv("AMQP_URL"), os.Getenv("NOTI_QUEUE_NAME"))

	routes.SetupNotificationRoutes(router)

	fmt.Println("Listening on port " + os.Getenv("PORT"))
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
