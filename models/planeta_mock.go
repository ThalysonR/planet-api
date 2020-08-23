package models

type PlanetaMock struct {
	BuscaPlanetaPorIDMock     func(ID string) (*Planeta, error)
	BuscaPlanetasPaginadoMock func(campoNome *string, campoValor *string, skip int, limit int) (*PaginaResultado, error)
	InserirPlanetaMock        func(input PlanetaInput) (*InsertResult, error)
	RemoverPlanetaMock        func(ID string) error
}

func (pm *PlanetaMock) BuscaPlanetaPorID(ID string) (*Planeta, error) {
	return pm.BuscaPlanetaPorIDMock(ID)
}

func (pm *PlanetaMock) BuscaPlanetasPaginado(campoNome *string, campoValor *string, skip int, limit int) (*PaginaResultado, error) {
	return pm.BuscaPlanetasPaginadoMock(campoNome, campoValor, skip, limit)
}

func (pm *PlanetaMock) InserirPlaneta(input PlanetaInput) (*InsertResult, error) {
	return pm.InserirPlanetaMock(input)
}

func (pm *PlanetaMock) RemoverPlaneta(ID string) error {
	return pm.RemoverPlanetaMock(ID)
}
