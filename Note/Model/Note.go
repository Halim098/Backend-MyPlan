package Model

import (
	DB "MyPlan-Note/Database"
	"crypto/sha1"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Content struct {
	Title   string `json:"title" bson:"title" binding:"required"`
	Content string `json:"content" bson:"content" binding:"required"`
}

type Note struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string            `json:"username" bson:"username" binding:"required"`
	Title   string             `json:"title" bson:"title" binding:"required"`
	Status string             `json:"status" bson:"status"`
	Link string               `json:"link" bson:"link"`
	Content []Content          `json:"content" bson:"content" binding:"required"`
}

func GetCollection() *mongo.Collection {
	return DB.Client.Database("myplan").Collection("notes")
}

func (n *Note) Save() (*mongo.InsertOneResult,error) {
	Collection := GetCollection()
	n.CreateLink()
	Data := bson.M{
		"title": n.Title,
		"content": n.Content,
		"username": n.Username,
		"link" : n.Link,
		"status": n.Status,
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
			"status": n.Status,
		},
	}
	return Collection.UpdateOne(DB.Ctx, filter, update)
}

func Find(username string) ([]Note, error) {
	Collection := GetCollection()
	cursor, err := Collection.Find(DB.Ctx, bson.M{"username": username})
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

func (n *Note) CreateLink() {
	var sha = sha1.New()
    sha.Write([]byte(n.ID.Hex()))
    var encrypted = sha.Sum(nil)
    n.Link = fmt.Sprintf("%x", encrypted)
}

func Delete(id string) (*mongo.DeleteResult, error) {
	Collection := GetCollection()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	return Collection.DeleteOne(DB.Ctx, filter)
}