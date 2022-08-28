package planets

import (
	"github.com/leonardoguarilha/domain/entities"
	"github.com/leonardoguarilha/repository"
	"strconv"
)

type GetPlanetByIdHandler struct {
	PlanetRepository repository.PlanetRepository
}

func (g *GetPlanetByIdHandler) Handle(id string) (*entities.PlanetApiReturn, error) {

	planetId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	planet, err := g.PlanetRepository.GetById(planetId)

	return planet, err
}
