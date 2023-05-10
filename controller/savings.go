package controller

import (
	"mini_project/model"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllSavings(c echo.Context) error {
	savings, err := usecase.GetAllSavings()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error fetching savings progress data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"savings": savings,
	})
}

func GetSavingByID(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	savings, err := usecase.GetSavingsByID(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error fetching savings data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"savingss": savings,
	})
}

func CreateSavings(c echo.Context) error {
	s := model.Savings{}
	c.Bind(&s)

	if err := usecase.CreateSavings(&s); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error creating savings data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"savings": "s",
	})
}

func UpdateSavings(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	s := model.Savings{}
	c.Bind(&s)
	s.ID = uint(id)

	if err := usecase.UpdateSavings(&s); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error updating savings data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"savings": s,
	})
}

func DeleteSavings(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeleteSavings(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error deleting savings data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
