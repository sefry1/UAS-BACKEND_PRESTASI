package main

import (
	"log"

	"prestasi_backend/app/config"
	"prestasi_backend/app/database"
	"prestasi_backend/app/route"
	"prestasi_backend/app/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadEnv()

	// Connect PostgreSQL
	pg, err := database.ConnectPostgre()
	if err != nil {
		log.Fatal(err)
	}
	database.PostgresDB = pg

	// Connect MongoDB
	mongo, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}
	database.MongoDB = mongo

	// 🔥 PENTING — Inisialisasi semua repository
	service.InitService()

	// Start Fiber
	app := fiber.New()
	route.SetupRoutes(app)

	port := config.Get("APP_PORT")
	log.Println("🚀 Server running on port", port)

	app.Listen(":" + port)
}
