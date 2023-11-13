package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RecipientID string             `json:"recipient_id" bson:"recipient_id"`
	OrderID     string             `json:"order_id" bson:"order_id"`
	IsRead      bool               `json:"is_read" bson:"is_read"`
	Type        string             `json:"type" bson:"type"`
	Timestamp   primitive.DateTime `json:"timestamp" bson:"timestamp"`
}
