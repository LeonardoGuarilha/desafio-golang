package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoguarilha/application/useCases/planets"
	"github.com/leonardoguarilha/domain/entities"
	"github.com/leonardoguarilha/repository"
	"net/http"
)

type PlanetController struct {
	PlanetService          repository.PlanetRepository
	CreatePlanetUseCase    planets.CreatePlanetUseCase
	GetAllPlanetsUseCase   planets.GetAllPlanetsUseCase
	GetPlanetByNameHandler planets.GetPlanetByNameHandler
	DeletePlanetHandler    planets.DeletePlanetHandler
	GetPlanetByIdHandler   planets.GetPlanetByIdHandler
}

func New(planetService repository.PlanetRepository,
	createPlanetUseCase planets.CreatePlanetUseCase,
	getAllPlanetsUseCase planets.GetAllPlanetsUseCase,
	getPlanetByNameHandler planets.GetPlanetByNameHandler,
	deletePlanetHandler planets.DeletePlanetHandler,
	getPlanetByIdHandler planets.GetPlanetByIdHandler) PlanetController {
	return PlanetController{
		PlanetService:          planetService,
		CreatePlanetUseCase:    createPlanetUseCase,
		GetAllPlanetsUseCase:   getAllPlanetsUseCase,
		GetPlanetByNameHandler: getPlanetByNameHandler,
		DeletePlanetHandler:    deletePlanetHandler,
		GetPlanetByIdHandler:   getPlanetByIdHandler,
	}
}

func (pc *PlanetController) CreatePlanet(context *gin.Context) {

	var planet entities.Planet

	if err := context.ShouldBindJSON(&planet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newPlanet, err := pc.CreatePlanetUseCase.Handle(planet)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, newPlanet)
}

func (pc *PlanetController) GetByName(context *gin.Context) {
	var planetName string = context.Param("name")

	planet, err := pc.GetPlanetByNameHandler.Handle(planetName)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Nenhum planeta com esse nome foi encontrado na base de dados"})
		return
	}

	context.JSON(http.StatusOK, planet)
}

func (pc *PlanetController) GetAll(context *gin.Context) {

	planetsFound, err := pc.GetAllPlanetsUseCase.Handle()

	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Nenhum planeta cadastrado na base de dados"})
		return
	}

	context.JSON(http.StatusOK, planetsFound)
}

func (pc *PlanetController) GetById(context *gin.Context) {
	var planetIdParam string = context.Param("id")

	planet, err := pc.GetPlanetByIdHandler.Handle(planetIdParam)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Nenhum planeta encontrado na base de dados com esse id"})
		return
	}

	context.JSON(http.StatusOK, planet)
}

func (pc *PlanetController) DeletePlanet(context *gin.Context) {
	var planetName string = context.Param("name")

	err := pc.DeletePlanetHandler.Handle(planetName)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Não existe planeta cadastrado com esse nome"})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "Planeta deletado com sucesso!"})
}

func (pc *PlanetController) RegisterPlanetRoutes(routerGroup *gin.RouterGroup) {
	planetroute := routerGroup.Group("/planet")

	planetroute.POST("/create", pc.CreatePlanet)
	planetroute.GET("/get/:name", pc.GetByName)
	planetroute.GET("/getall", pc.GetAll)
	planetroute.GET("/getbyid/:id", pc.GetById)
	planetroute.DELETE("/delete/:name", pc.DeletePlanet)
}
