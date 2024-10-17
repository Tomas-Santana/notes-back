package main

import (
	"log"
	"notes-back/api"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"flag"
	"notes-back/database"
	"github.com/resend/resend-go/v2"
)

func main() {
	env := flag.String("environment", "local", "environment") 
	listenaddr := flag.String("listenaddr", "192.168.18.12:8080", "server listen address")
	flag.Parse()

	
	loadEnv(*env)
	mongo_uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("DB_NAME")
	resend_api_key := os.Getenv("RESEND_API_KEY")

	if resend_api_key == "" {
		panic("RESEND_API_KEY must be set in .env file")
	}

	if mongo_uri == "" || db_name == "" {
		panic("MONGO_URI and DB_NAME must be set in .env file")
	}

	emailClient := resend.NewClient(resend_api_key)

	var validator = validator.New()

	db := database.NewMongoDatabase(mongo_uri, db_name)
	if err := db.Connect(); err != nil {
		panic(err)
	}
	server := api.NewServer(*listenaddr, db, validator, emailClient)

	log.Fatal(server.Start())
}


func loadEnv(env string) {
	err := godotenv.Load(".env." + env)
	if err != nil {
		panic(err)
	}
}
