package controller

import (
	"mini_project/model"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetKasByIdController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}
	kas, e := usecase.GetKasById(uint(id))

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"kas":    kas,
	})
}

func CreateKasController(c echo.Context) error {
	kas := model.Kas{}
	c.Bind(&kas)

	if err := usecase.CreateKas(&kas); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create kas",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new kas",
		"kas":     kas,
	})
}

func UpdateKasController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	kas := model.Kas{}
	c.Bind(&kas)
	kas.ID = uint(id)
	if err := usecase.UpdateKas(&kas); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update kas",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf kas tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update kas",
	})
}
