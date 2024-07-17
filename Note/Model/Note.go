package Model

import (
	DB "MyPlan-Note/Database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Content struct {
	Title   string `json:"content_title" bson:"content_title"`
	Content string `json:"content" bson:"content"`
}

type Note struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string             `json:"title" bson:"title"`
	Content []Content          `json:"content" bson:"content"`
}

func GetCollection() *mongo.Collection {
	return DB.Client.Database("myplan").Collection("notes")
}

func (n *Note) Save() (*mongo.InsertOneResult,error) {
	Collection := GetCollection()
	Data := bson.M{
		"title": n.Title,
		"content": n.Content,
	}
	return Collection.InsertOne(DB.Ctx, Data)
}