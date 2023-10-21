package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dineflow-project/dineflow-notification-service/database"
	"github.com/dineflow-project/dineflow-notification-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNotificationsHandler(c *gin.Context) {
	recipientID := c.Param("recipientID")
	collection := database.Client.Database("noti-service").Collection("notifications")

	cursor, err := collection.Find(context.TODO(), bson.M{"recipient_id": recipientID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}
	defer cursor.Close(context.TODO())

	var notifications []models.Notification
	for cursor.Next(context.TODO()) {
		var notification models.Notification
		if err := cursor.Decode(&notification); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode notification"})
			return
		}
		notifications = append(notifications, notification)
	}

	// Update is_read of the notifications to true
	_, err = collection.UpdateMany(
		context.TODO(),
		bson.M{"recipient_id": recipientID, "is_read": false},
		bson.M{"$set": bson.M{"is_read": true}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

func GetUnreadNotificationsCountHandler(c *gin.Context) {
	recipientID := c.Param("recipientID")
	collection := database.Client.Database("noti-service").Collection("notifications")

	count, err := collection.CountDocuments(context.TODO(), bson.M{
		"recipient_id": recipientID,
		"is_read":      false,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func UpdateNotifications(msg string) error {
	var notification models.Notification
	err := json.Unmarshal([]byte(msg), &notification)
	if err != nil {
		return err
	}

	collection := database.Client.Database(os.Getenv("MONGO_DATABASE")).Collection("notifications")

	if notification.Type == "delete" {
		_, err = collection.DeleteMany(context.Background(), bson.M{"order_id": notification.OrderID})
		if err != nil {
			return err
		}
	} else {
		_, err = collection.InsertOne(context.Background(), notification)
		if err != nil {
			return err
		}
	}

	return nil
}
