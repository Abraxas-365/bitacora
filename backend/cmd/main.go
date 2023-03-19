package main

import (
	"fmt"
	"github/Abraxas-365/bitacora/internal/mongo"
	"github/Abraxas-365/bitacora/pkg/auth"
	"github/Abraxas-365/bitacora/pkg/report"
	"github/Abraxas-365/bitacora/pkg/user"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	mongoUri := os.Getenv("MONGODB_URI")
	mongoClient, _ := mongo.MongoFactory(mongoUri, "Bitacora", 10)
	elasticClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://es-container:9200"},
		// Username:  "elastic",
		// Password:  "elstic",
	})
	if err != nil {
		log.Panic(err)
	}
	// Test the connection to Elasticsearch
	res, err := elasticClient.Info()
	if err != nil {
		log.Fatalf("Error getting Elasticsearch info: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	fmt.Println(res.String())

	report.ModuleFactory(app, mongoClient, elasticClient)
	userApp := user.ModuleFactory(app, mongoClient)
	auth.ModuleFactory(app, userApp)

	app.Listen(":1234")

}
