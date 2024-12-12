package space

import (
	"stellaris/unit/common"
	"stellaris/unit/resource"
)

type CelestialType int

const (
	CelestialInvalid = CelestialType(iota)
	CelestialPlanet
	CelestialStar
	CelestialFairway
	CelestialWormhole
)

type Celestial struct {
	Name string

	Pos common.GalaxyPosition
}

func (Celestial) Type() CelestialType {
	return CelestialInvalid
}

type Capturable struct {
	Owner  int
	Occupy int
}

type Planet struct {
	Celestial
	Capturable
	Resources []resource.Resource
}

func (Planet) Type() CelestialType {
	return CelestialPlanet
}

type Star struct {
	Celestial
	Capturable
	Resources []resource.Resource
}

func (Star) Type() CelestialType {
	return CelestialStar
}

type Wormhole struct {
	Celestial
	Capturable
}

func (Wormhole) Type() CelestialType {
	return CelestialWormhole
}

type Fairway struct {
	Celestial
}

func (Fairway) Type() CelestialType {
	return CelestialFairway
}
