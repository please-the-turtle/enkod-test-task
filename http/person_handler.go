package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/please-the-turtle/encod-test-task/app"
)

type ResponseError struct {
	Message string `json:"message"`
}

type PersonHandler struct {
	PLogic app.PersonLogic
}

func NewPersonHandler(e *echo.Echo, l app.PersonLogic) {
	handler := &PersonHandler{
		PLogic: l,
	}
	e.GET("/person", handler.FetchPersons)
	e.POST("/person", handler.Create)
	e.GET("/person/:id", handler.GetByID)
	e.PUT("/person/:id", handler.Update)
	e.DELETE("/person/:id", handler.Delete)
}

func (p *PersonHandler) FetchPersons(c echo.Context) error {
	ctx := c.Request().Context()
	persons, err := p.PLogic.Fetch(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, persons)
}

func (p *PersonHandler) Create(c echo.Context) error {
	return errors.New("method not implemented")
}

func (p *PersonHandler) GetByID(c echo.Context) error {
	return errors.New("method not implemented")
}

func (p *PersonHandler) Update(c echo.Context) error {
	return errors.New("method not implemented")
}

func (p *PersonHandler) Delete(c echo.Context) error {
	return errors.New("method not implemented")
}
