package main

import (
	"MyPlan-Note/Database"

	//env
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file")
	}

	client, ctx, cancel, err := Database.Connect()
    if err != nil {
        panic(err)
    }

    defer Database.Close(client, ctx, cancel)

    Database.Ping(client, ctx)
}