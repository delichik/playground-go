package resource

type Type int

const (
	Energy = Type(iota)
	Mineral
	Food
	Luxury
	Alloy

	RareCrystal
	HeterostarGas
	ExplosiveParticle

	Physics
	Engineering
	Sociology
)

type Resource struct {
	Type  Type
	Count int64
}
