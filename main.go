package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/programmercintasunnah/dompet-keluarga/config"
	"github.com/programmercintasunnah/dompet-keluarga/routes"
)

func main() {
	fmt.Println("Ini adalah aplikasi sederhana (dompet keluarga)")
	fmt.Println("di create pada 1 januari 2025")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	routes.SetupRoutes(app)

	port := config.GetEnv("PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}
