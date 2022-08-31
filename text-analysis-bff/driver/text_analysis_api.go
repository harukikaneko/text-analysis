package driver

import (
	"echo-get-started/config"
	"echo-get-started/port"
	"echo-get-started/types"
	"encoding/json"
	"fmt"
	"net/url"
)

type TextAnalysisApi struct{}

func NewNounPort() port.NounPort {
	port := &TextAnalysisApi{}
	return port
}

func (*TextAnalysisApi) GetCountsByNoun(query string) ([]types.CountsByNoun, error) {
	uri := fmt.Sprintf("%s/nouns", config.GetConfig().TextAnalysisEndpoint)

	query_params := make(url.Values)
	query_params.Set("query", query)

	byte, err := doGetWithQuery(uri, query_params)
	if err != nil {
		return []types.CountsByNoun{}, err
	}

	var value []types.CountsByNoun
	if err := json.Unmarshal(byte, &value); err != nil {
		return []types.CountsByNoun{}, err
	}

	return value, nil
}
