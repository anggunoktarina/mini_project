package route

import (
	"mini_project/controller"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controller.GetUserController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.PUT("/:id", controller.UpdateUserController)
	user.DELETE("/:id", controller.DeleteUserController)

	//income collection
	income := e.Group("/incomes")
	income.GET("", controller.GetIncomeController)
	income.POST("", controller.CreateIncomeController)
	income.PUT("/:id", controller.UpdateIncomeController)
	income.DELETE("/:id", controller.DeleteIncomeController)

	//expense collection
	expense := e.Group("/expenses")
	expense.GET("", controller.GetExpenseController)
	expense.POST("", controller.CreateExpenseController)
	expense.PUT("/:id", controller.UpdateExpenseController)
	expense.DELETE("/:id", controller.DeleteExpenseController)

	//kas collection
	kas := e.Group("/kas")
	kas.GET("/:id", controller.GetKasByIdController)
	kas.GET("/kas/:id/", controller.GetRemainingBalanceByKasID)
	kas.POST("", controller.CreateKasController)
	kas.PUT("/:id", controller.UpdateKasController)

	//saving progress
	savings := e.Group("/savings")
	savings.GET("", controller.GetAllSavings)
	savings.GET("/:id", controller.GetSavingByID)
	savings.POST("", controller.CreateSavings)
	savings.PUT("/:id", controller.UpdateSavings)
	savings.DELETE("/:id", controller.DeleteSavings)

}
