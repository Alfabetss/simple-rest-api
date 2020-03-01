package controller

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Alfabetss/simple-rest-api/config"
	"github.com/Alfabetss/simple-rest-api/service"
	"github.com/labstack/echo"
)

// TalentController exported
type TalentController struct {
	cfg config.Configuration
	db  *sql.DB
}

// NewTalentController constructor
func NewTalentController(
	config config.Configuration,
	db *sql.DB,
) TalentController {
	return TalentController{
		cfg: config,
		db:  db,
	}
}

// CreateTalent handler for create talent
func (t TalentController) CreateTalent(e echo.Context) error {
	request := new(service.CreateTalentRequest)
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "bad request")
	}

	tService := service.NewTalentServiceImpl(t.db)
	err := tService.CreateTalent(e.Request().Context(), request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, "success create new user")
}

// FindTalent handler for find talent
func (t TalentController) FindTalent(e echo.Context) error {
	log.Printf("config : %+v", t.cfg)
	param := e.Param("talentID")
	talentID, err := strconv.ParseInt(param, 10, 64)
	if err != nil || talentID == 0 {
		return e.JSON(http.StatusBadRequest, "bad request")
	}

	tService := service.NewTalentServiceImpl(t.db)
	resp, err := tService.FindTalent(e.Request().Context(), talentID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, resp)
}

// Delete handler for delete talent
func (t TalentController) Delete(e echo.Context) error {
	param := e.Param("talentID")
	talentID, err := strconv.ParseInt(param, 10, 64)
	if err != nil || talentID == 0 {
		return e.JSON(http.StatusBadRequest, "bad request")
	}

	tService := service.NewTalentServiceImpl(t.db)
	err = tService.Delete(e.Request().Context(), talentID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, "success delete talent")
}

// UpdateTalent handler update talent data
func (t TalentController) UpdateTalent(e echo.Context) error {
	request := new(service.UpdateTalentRequest)
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request")
	}

	tService := service.NewTalentServiceImpl(t.db)
	err := tService.UpdateTalent(e.Request().Context(), *request)
	if err != nil {
		if err == sql.ErrNoRows {
			return e.JSON(http.StatusNotFound, "data not found")
		}
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, "success update talent")
}

// UpdateTalentExperience handler update talent data
func (t TalentController) UpdateTalentExperience(e echo.Context) error {
	request := new(service.UpdateExperienceRequest)
	if err := e.Bind(request); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request")
	}

	tService := service.NewTalentServiceImpl(t.db)
	err := tService.UpdateExperience(e.Request().Context(), *request)
	if err != nil {
		if err == sql.ErrNoRows {
			return e.JSON(http.StatusNotFound, "data not found")
		}
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, "success update talent experience")
}
