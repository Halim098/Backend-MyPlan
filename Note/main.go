package main

import (
	"MyPlan-Note/Database"
	"MyPlan-Note/Router"

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

    err = Database.Ping()
	if err != nil {
		panic(err)
	}

	r := Router.SetupRouter()
	r.Run(":8081")
}