package main

import (
	"MyPlant-User/Database"
	"MyPlant-User/Model"
	"MyPlant-User/Router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file")
	}

	Database.Connection()
	Database.DB.AutoMigrate(&Model.User{})

	r := Router.SetupRouter()
	r.Run(":8080")
}

