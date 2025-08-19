package main

import (
	"log"
	"svi-be/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Env tidak ditemukan")
	}

	r := router.SetupRouter()

	log.Println("Running on :8081")
	r.Run(":8081")
}
