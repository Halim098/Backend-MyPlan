package Database

import (
	// import mongodb driver
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc){

    defer cancel()

    defer func(){
        if err := client.Disconnect(ctx); err != nil{
            panic(err)
        }
    }()
}

