package controller

import (
	"mini_project/model"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetExpenseController(c echo.Context) error {
	expenses, err := usecase.GetExpense()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to get expenses",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"expenses": expenses,
	})
}

func CreateExpenseController(c echo.Context) error {
	expense := model.Expense{}
	c.Bind(&expense)

	err := usecase.CreateExpense(&expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create expense",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "Expense created successfully",
		"expense": expense,
	})
}

func UpdateExpenseController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	expense := model.Expense{}
	c.Bind(&expense)
	expense.ID = uint(id)

	err := usecase.UpdateExpense(&expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update expense",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Expense updated successfully",
		"expense": expense,
	})
}

func DeleteExpenseController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	err := usecase.DeleteExpense(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to delete expense",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Expense deleted successfully",
	})
}
