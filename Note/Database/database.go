package Database

import (
	// import mongodb driver
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var Ctx context.Context

func Close(client *mongo.Client, cancel context.CancelFunc){

    defer cancel()

    defer func(){
        if err := client.Disconnect(Ctx); err != nil{
            panic(err)
        }
    }()
}

func Connect()(context.CancelFunc, error) {
	var err error
	var cancel context.CancelFunc

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
    Ctx, cancel = context.WithTimeout(context.Background(), 30 * time.Second)

    Client, err = mongo.Connect(Ctx, options.Client().ApplyURI(uri))
    return cancel, err
}

func Ping() error {
	collection := Client.Database("myplan").Collection("notes")

	fmt.Println(collection)

    if err := Client.Ping(Ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}