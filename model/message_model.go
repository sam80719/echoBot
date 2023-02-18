package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"user_id"`
	Message     string             `bson:"message"`
	MessageType string             `bson:"message_type"`
	SendTime    primitive.DateTime `bson:"send_time"`
}
