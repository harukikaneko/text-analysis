package use_case_test

import (
	"echo-get-started/types"
	"echo-get-started/use_case"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCountsByNoun(t *testing.T) {
	nounPort := new(MockNounPort)

	target := use_case.NewNounUseCase(nounPort)

	returnValue := []types.CountsByNoun{
		{
			Noun:   "hoge",
			Counts: 2,
		},
		{
			Noun:   "fuga",
			Counts: 1,
		},
	}

	expected :=
		[]types.CountsByNoun{
			{
				Noun:   "hoge",
				Counts: 2,
			},
			{
				Noun:   "fuga",
				Counts: 1,
			},
		}

	nounPort.On("GetCountsByNoun", "hogefuga").Return(returnValue, nil)
	actual, _ := target.GetCountsByNoun("hogefuga")

	assert.Equal(t, actual, expected)
}

type MockNounPort struct {
	mock.Mock
}

func (m MockNounPort) GetCountsByNoun(query string) ([]types.CountsByNoun, error) {
	args := m.Called(query)
	return args.Get(0).([]types.CountsByNoun), args.Error(1)
}
