package entities

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

type Planet struct {
	PlanetApiId int
	Name        string `json:"name" bson:"name" valid:"notnull"`
	Climate     string `json:"climate" bson:"climate" valid:"notnull"`
	Terrain     string `json:"terrain" bson:"terrain" valid:"notnull"`
}

type PlanetApiReturn struct {
	PlanetApiId int
	Name        string   `json:"name" bson:"name"`
	Climate     string   `json:"climate" bson:"climate"`
	Terrain     string   `json:"terrain" bson:"terrain"`
	Films       []string `json:"films" bson:"films"`
}

func (p *Planet) IsValid(planet *Planet) error {
	_, err := govalidator.ValidateStruct(p)

	if planet.Name == "" || planet.Climate == "" || planet.Terrain == "" {
		return errors.New("Todos os campos devem estar preenchidos para criar um novo planeta")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPlanet(planetId int, name string, climate string, terrain string) (*Planet, error) {
	planet := Planet{
		PlanetApiId: planetId,
		Name:        name,
		Climate:     climate,
		Terrain:     terrain,
	}

	err := planet.IsValid(&planet)
	if err != nil {
		return nil, err
	}

	return &planet, nil
}
