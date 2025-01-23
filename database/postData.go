package database

import "go.mongodb.org/mongo-driver/mongo"

func CreateTask(data interface{}, collection *mongo.Collection) {

	collection.InsertOne(ctx, data)

}
