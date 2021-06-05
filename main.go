package main

import (
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	database := connectToDatabase()
	defer database.close()

	createBot(database)
}
