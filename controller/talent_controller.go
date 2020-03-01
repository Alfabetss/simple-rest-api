package controller

import (
	"net/http"
	"strconv"

	"github.com/Alfabetss/simple-rest-api/service"
	"github.com/labstack/echo"
)

// CreateTalent handler for create talent
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

// FindTalent handler for find talent
func FindTalent(e echo.Context) error {
	param := e.Param("talentID")
	talentID, err := strconv.ParseInt(param, 10, 64)
	if err != nil || talentID == 0 {
		return e.JSON(http.StatusBadRequest, "bad request")
	}

	tService := service.NewTalentServiceImpl()
	resp, err := tService.FindTalent(e.Request().Context(), talentID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, resp)
}
