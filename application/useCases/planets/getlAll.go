package planets

import (
	"github.com/leonardoguarilha/domain/entities"
	"github.com/leonardoguarilha/repository"
)

type GetAllPlanetsUseCase struct {
	PlanetRepository repository.PlanetRepository
}

func (g *GetAllPlanetsUseCase) Handle() ([]*entities.Planet, error) {
	planets, err := g.PlanetRepository.GetAll()

	return planets, err
}
