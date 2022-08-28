package repository

import (
	"github.com/leonardoguarilha/domain/entities"
)

type PlanetRepository interface {
	CreatePlanet(planet *entities.Planet) (*entities.Planet, error)
	GetByName(*string) (*entities.PlanetApiReturn, error)
	GetById(int) (*entities.PlanetApiReturn, error)
	GetAll() ([]*entities.Planet, error)
	DeletePlanet(*string) error
	GetTerrainsWithMoviesById(id int) entities.PlanetApiReturn
	GetCountPlanets() int
}
