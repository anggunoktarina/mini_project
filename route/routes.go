package route

import (
	"mini_project/controller"
	"mini_project/middlewares"
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

	middlewares := middlewares.NewMiddlewares()
	e.Use(middlewares.Logmiddleware)
	e.Pre(middlewares.RemoveTrailingSlash())

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/register", controller.CreateUserController)
	e.POST("/login", controller.LoginUserController)

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
	expense.POST("", controller.CreateIncomeController)
	expense.PUT("/:id", controller.UpdateExpenseController)
	expense.DELETE("/:id", controller.DeleteExpenseController)

	//kas collection
	kas := e.Group("/kas")
	kas.GET("", controller.GetKasByIdController)
	kas.POST("", controller.CreateKasController)
	kas.PUT("/:id", controller.UpdateKasController)

	//template savings collection
	template_savings := e.Group("/template_savings")
	template_savings.GET("", controller.GetTemplateSavings)
	template_savings.POST("", controller.CreateTemplateSavings)
	template_savings.PUT("/:id", controller.UpdateTemplateSavingsController)

	//saving progress
	savings_progress := e.Group("/savings_progress")
	savings_progress.GET("", controller.GetSavingsProgress)
	savings_progress.GET("", controller.GetSavingsProgressByID)
	savings_progress.POST("", controller.CreateSavingsProgress)
	savings_progress.PUT("/:id", controller.UpdateSavingsProgress)
	savings_progress.DELETE("/:id", controller.DeleteSavingsProgress)

}
