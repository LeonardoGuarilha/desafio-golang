package planets

import "github.com/leonardoguarilha/repository"

type DeletePlanetHandler struct {
	PlanetRepository repository.PlanetRepository
}

func (d *DeletePlanetHandler) Handle(name string) error {
	err := d.PlanetRepository.DeletePlanet(&name)

	return err
}
