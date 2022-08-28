package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/leonardoguarilha/application/useCases/planets"
	"github.com/leonardoguarilha/configurations/database"
	"github.com/leonardoguarilha/controllers"
	"github.com/leonardoguarilha/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	ctx         context.Context
	mongoClient *mongo.Client
)

func main() {
	planetCollection := database.ConnectDataBase()
	planetRepository := repository.NewPlanetService(planetCollection, ctx)

	createPlanetUseCase := planets.CreatePlanetUseCase{PlanetRepository: planetRepository}
	getAllPlanetUseCase := planets.GetAllPlanetsUseCase{PlanetRepository: planetRepository}
	getPlanetByNameHandler := planets.GetPlanetByNameHandler{PlanetRepository: planetRepository}
	deletePlanetHandler := planets.DeletePlanetHandler{PlanetRepository: planetRepository}
	getPlanetByIdHandler := planets.GetPlanetByIdHandler{PlanetRepository: planetRepository}
	planetController := controllers.New(planetRepository, createPlanetUseCase, getAllPlanetUseCase,
		getPlanetByNameHandler, deletePlanetHandler, getPlanetByIdHandler)

	server := gin.Default()

	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")
	planetController.RegisterPlanetRoutes(basepath)

	log.Fatal(server.Run(":8080"))
}
