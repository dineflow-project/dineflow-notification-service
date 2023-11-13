package routes

import (
	"github.com/dineflow-project/dineflow-notification-service/controllers"
	"github.com/gin-gonic/gin"
)

func SetupNotificationRoutes(router *gin.Engine) {
	// Define routes for notifications
	router.GET("/notifications/:recipientID", controllers.GetNotificationsHandler)
	router.GET("/notifications/unread/:recipientID", controllers.GetUnreadNotificationsCountHandler)
}
