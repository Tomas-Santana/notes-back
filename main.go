package main

import (
	"log"
	"notes-back/api"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"notes-back/database"
)

func main() {
	loadEnv()
	mongo_uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("DB_NAME")

	if mongo_uri == "" || db_name == "" {
		panic("MONGO_URI and DB_NAME must be set in .env file")
	}

	var validator = validator.New()

	db := database.NewMongoDatabase(mongo_uri, db_name)
	if err := db.Connect(); err != nil {
		panic(err)
	}
	server := api.NewServer("localhost:8080", db, validator)

	log.Fatal(server.Start())
}


func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
