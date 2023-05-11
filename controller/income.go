package controller

import (
	"mini_project/model"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetIncomeController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	income, err := usecase.GetIncome(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get income",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"income": income,
	})
}

func CreateIncomeController(c echo.Context) error {
	income := model.Income{}
	c.Bind(&income)

	if err := usecase.CreateIncome(&income); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create income",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new income",
		"income":  income,
	})
}

func DeleteIncomeController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteIncome(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete income",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf income tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete income",
	})
}

func UpdateIncomeController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	income := model.Income{}
	c.Bind(&income)
	income.ID = uint(id)
	if err := usecase.UpdateIncome(&income); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update income",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf income tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update income",
	})
}
