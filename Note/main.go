package main

import (
	"MyPlan-Note/Database"
)

func main() {
	client, ctx, cancel, err := Database.Connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }

    defer Database.Close(client, ctx, cancel)

    Database.Ping(client, ctx)
}