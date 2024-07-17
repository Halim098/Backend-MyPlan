package main

import (
	"MyPlan-Note/Database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file")
	}

	cancel, err := Database.Connect()
    if err != nil {
        panic(err)
    }

    defer Database.Close(Database.Client, cancel)

    Database.Ping()
}