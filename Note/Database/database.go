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

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc){

    defer cancel()

    defer func(){
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
}

func Connect()(*mongo.Client, context.Context, context.CancelFunc, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
    ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}

func Ping(client *mongo.Client, ctx context.Context) error {

    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}