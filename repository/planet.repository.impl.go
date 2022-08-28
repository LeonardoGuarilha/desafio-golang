package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/leonardoguarilha/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
)

type PlanetRepositoryImpl struct {
	planetCollection *mongo.Collection
	context          context.Context
}

func NewPlanetService(planetCollection *mongo.Collection, context context.Context) PlanetRepository {
	return &PlanetRepositoryImpl{
		planetCollection: planetCollection,
		context:          context,
	}
}

func (p *PlanetRepositoryImpl) CreatePlanet(planet *entities.Planet) (*entities.Planet, error) {
	planetsCount := p.GetCountPlanets()
	planet.PlanetApiId = planetsCount + 1

	newPlanet, err := entities.NewPlanet(planet.PlanetApiId, planet.Name, planet.Climate, planet.Terrain)

	if err != nil {
		return nil, err
	}

	_, err = p.planetCollection.InsertOne(p.context, newPlanet)
	return newPlanet, err
}

func (p *PlanetRepositoryImpl) GetByName(name *string) (*entities.PlanetApiReturn, error) {
	var planet *entities.PlanetApiReturn

	query := bson.D{bson.E{Key: "name", Value: name}}

	err := p.planetCollection.FindOne(p.context, query).Decode(&planet)

	if err != nil {
		return nil, err
	}

	terrain := p.GetTerrainsWithMoviesById(planet.PlanetApiId)

	planet.Films = terrain.Films

	return planet, err
}

func (p *PlanetRepositoryImpl) GetTerrainsWithMoviesById(id int) entities.PlanetApiReturn {
	requestURL := fmt.Sprintf("http://swapi.dev/api/planets/%d", id)

	response, err := http.Get(requestURL)

	if err != nil {
		fmt.Println("Erro ao carregar planetas")
	}

	data, _ := ioutil.ReadAll(response.Body)

	var planets entities.PlanetApiReturn

	json.Unmarshal(data, &planets)

	return planets
}

func (p *PlanetRepositoryImpl) GetAll() ([]*entities.Planet, error) {
	var planets []*entities.Planet
	cursor, err := p.planetCollection.Find(p.context, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(p.context) {
		var planet entities.Planet
		err := cursor.Decode(&planet)
		if err != nil {
			return nil, err
		}
		planets = append(planets, &planet)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(p.context)

	if len(planets) == 0 {
		return nil, err
	}

	return planets, nil
}

func (p *PlanetRepositoryImpl) GetById(id int) (*entities.PlanetApiReturn, error) {
	var planet *entities.PlanetApiReturn

	query := bson.D{bson.E{Key: "planetapiid", Value: id}}

	err := p.planetCollection.FindOne(p.context, query).Decode(&planet)

	if err != nil {
		return nil, err
	}

	terrain := p.GetTerrainsWithMoviesById(planet.PlanetApiId)

	planet.Films = terrain.Films

	return planet, err
}

func (p *PlanetRepositoryImpl) DeletePlanet(name *string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result, _ := p.planetCollection.DeleteOne(p.context, filter)
	if result.DeletedCount != 1 {
		return errors.New("Não foi encontrado nenhum planeta com esse nome.")
	}

	return nil
}

func (p *PlanetRepositoryImpl) GetCountPlanets() int {
	var planets []*entities.Planet
	cursor, err := p.planetCollection.Find(p.context, bson.D{{}})
	if err != nil {
		return 0
	}
	for cursor.Next(p.context) {
		var planet entities.Planet
		err := cursor.Decode(&planet)
		if err != nil {
			return 0
		}
		planets = append(planets, &planet)
	}

	if err := cursor.Err(); err != nil {
		return 0
	}

	cursor.Close(p.context)

	return len(planets)
}
