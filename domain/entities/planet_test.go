package entities

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEntity_GivenValidInformation_WhenCallsNewPlanet_ShouldReturnANewPlanet(t *testing.T) {
	planetId := 1
	name := "Test Planet"
	climate := "Desert"
	terrain := "Arid"

	newPlanet, err := NewPlanet(planetId, name, climate, terrain)

	require.Nil(t, err)
	require.Equal(t, newPlanet.PlanetApiId, planetId)
	require.Equal(t, newPlanet.Name, name)
	require.Equal(t, newPlanet.Climate, climate)
	require.Equal(t, newPlanet.Terrain, terrain)
}

func TestEntity_GivenInvalidClimateAndTerrain_WhenCallsNewPlanet_ShouldReturnError(t *testing.T) {
	planetId := 1
	name := "Test Planet"
	climate := ""
	terrain := ""

	_, err := NewPlanet(planetId, name, climate, terrain)

	require.NotNil(t, err)
}
