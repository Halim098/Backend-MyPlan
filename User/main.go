package main

import (
	"MyPlan-User/Database"
	"MyPlan-User/Model"
	"MyPlan-User/Router"

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

