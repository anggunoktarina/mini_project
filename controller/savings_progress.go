package controller

import (
	"mini_project/model"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetSavingsProgress(c echo.Context) error {
	savingsProgress, err := usecase.GetSavingsProgress()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error fetching savings progress data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success",
		"savingsProgress": savingsProgress,
	})
}

func GetSavingsProgressByID(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	savingsProgress, err := usecase.GetSavingsProgress(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error fetching savings progress data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success",
		"savingsProgress": savingsProgress,
	})
}

func CreateSavingsProgress(c echo.Context) error {
	sp := model.Savings_Progress{}
	c.Bind(&sp)

	if err := usecase.CreateSavingsProgress(&sp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error creating savings progress data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success",
		"savingsProgress": sp,
	})
}

func UpdateSavingsProgress(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	sp := model.Savings_Progress{}
	c.Bind(&sp)
	sp.ID = uint(id)

	if err := usecase.UpdateSavingsProgress(&sp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error updating savings progress data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success",
		"savingsProgress": sp,
	})
}

func DeleteSavingsProgress(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeleteSavingsProgress(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error deleting savings progress data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
