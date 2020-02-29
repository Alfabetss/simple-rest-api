package controller

import (
	"net/http"

	"github.com/Alfabetss/simple-rest-api/service"
	"github.com/labstack/echo"
)

// CreateTalent handler
func CreateTalent(e echo.Context) error {
	request := new(service.CreateTalentRequest)
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "bad request")
	}

	tService := service.NewTalentServiceImpl()
	err := tService.CreateTalent(e.Request().Context(), request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, "success create new user")
}
