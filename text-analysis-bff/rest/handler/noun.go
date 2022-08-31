package handler

import (
	"echo-get-started/use_case"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NounHandler struct {
	useCase *use_case.NounUseCase
}

func NewNounHandler(useCase *use_case.NounUseCase) *NounHandler {
	return &NounHandler{useCase}
}

func (n *NounHandler) GetCountByNoun() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := c.QueryParams()
		countsByNoun, err := n.useCase.GetCountsByNoun(params.Get("query"))

		if err != nil {
			c.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, countsByNoun)
	}
}
