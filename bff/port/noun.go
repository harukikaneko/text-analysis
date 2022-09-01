package port

import "echo-get-started/types"

type NounPort interface {
	GetCountsByNoun(query string) ([]types.CountsByNoun, error)
}
