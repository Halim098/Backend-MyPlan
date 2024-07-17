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

func (n *Note) Update(id string) (*mongo.UpdateResult, error) {
	Collection := GetCollection()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"title": n.Title,
			"content": n.Content,
		},
	}
	return Collection.UpdateOne(DB.Ctx, filter, update)
}

func Find() ([]Note, error) {
	Collection := GetCollection()
	cursor, err := Collection.Find(DB.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(DB.Ctx)

	var notes []Note
	for cursor.Next(DB.Ctx) {
		var note Note
		if err := cursor.Decode(&note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func FindOne(id string) (Note, error) {
	Collection := GetCollection()
	var note Note
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return note, err
	}
	err = Collection.FindOne(DB.Ctx, bson.M{"_id": objectID}).Decode(&note)
	return note, err
}