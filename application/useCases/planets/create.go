package planets

import (
	"errors"
	"fmt"
	"github.com/leonardoguarilha/domain/entities"
	"github.com/leonardoguarilha/repository"
)

type CreatePlanetUseCase struct {
	PlanetRepository repository.PlanetRepository
}

func (c *CreatePlanetUseCase) Handle(planet entities.Planet) (*entities.Planet, error) {

	planetExists, _ := c.PlanetRepository.GetByName(&planet.Name)

	if planetExists != nil {
		message := fmt.Sprintf("Já existe um planeta cadastrado com esse nome, o mesmo possui o id: %d", planetExists.PlanetApiId)
		return nil, errors.New(message)
	}

	newPlanet, err := c.PlanetRepository.CreatePlanet(&planet)

	if err != nil {
		return nil, err
	}

	return newPlanet, nil
}
