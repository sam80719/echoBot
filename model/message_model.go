package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"reply_token"`
	Message     string             `bson:"message"`
	MessageType string             `bson:"type"`
	SendTime    primitive.DateTime `bson:"timestamp"`
	MessageData string             `bson:"message_data"`
}

func InsertMessageToMongo(db *mongo.Database, message *Message) error {
	_, err := db.Collection("messages").InsertOne(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}

func GetMessagesFromMongo(db *mongo.Database, limit int64) ([]Message, error) {
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	cur, err := db.Collection("messages").Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var messages []Message
	for cur.Next(context.Background()) {
		var message Message
		if err := cur.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
