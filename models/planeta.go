package models

type PlanetaDatasource interface{}

// Planeta representa um planeta
type Planeta struct {
	Nome    string
	Clima   string
	Terreno string
}
