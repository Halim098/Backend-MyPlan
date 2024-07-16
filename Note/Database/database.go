package Database

import (
	// import mongodb driver
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc){

    defer cancel()

    defer func(){
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
}

func connect(uri string)(*mongo.Client, context.Context, context.CancelFunc, error) {

    ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}


