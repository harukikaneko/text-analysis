package use_case

import (
	"echo-get-started/port"
	"echo-get-started/types"
)

type NounUseCase struct {
	port port.NounPort
}

func NewNounUseCase(port port.NounPort) *NounUseCase {
	return &NounUseCase{port}
}

func (m *NounUseCase) GetCountsByNoun(query string) ([]types.CountsByNoun, error) {
	countsByNoun, err := m.port.GetCountsByNoun(query)
	if err != nil {
		return []types.CountsByNoun{}, err
	}

	return countsByNoun, nil
}
