package planets

import (
	"github.com/leonardoguarilha/domain/entities"
	"github.com/leonardoguarilha/repository"
)

type GetPlanetByNameHandler struct {
	PlanetRepository repository.PlanetRepository
}

func (g *GetPlanetByNameHandler) Handle(planetName string) (*entities.PlanetApiReturn, error) {

	planet, err := g.PlanetRepository.GetByName(&planetName)

	if err != nil {
		return nil, err
	}

	return planet, nil
}
