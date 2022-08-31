package types

import "fmt"

type JSON = map[string]interface{}

type HttpError struct {
	StatusCode int
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("http error. status code: %d", e.StatusCode)
}

func (e *HttpError) GetStatusCode() int {
	return e.StatusCode
}

type StatusCodeHolder interface {
	GetStatusCode() int
}
