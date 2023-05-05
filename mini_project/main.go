package main

import (
	"mini_project/config"
	"mini_project/model"
	"mini_project/middleware"
	"mini_project/route"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)
func init() {
	InitDB()
	InitialMigration()
}

//database connection
func InitDB() {
	// declare struct config & variable connectionString
	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	ID int `json:"id" form "id"`
	Name string `json:"name" form "name"`
	Email string `json:"email" form "email'`
	Password string `json:"password" form "password"`
}

type Income struct {
	Income_ID int `json:"income_id" form "income_id"`
	Kas_ID int `json:"kas_id" form "kas_id"`
	Nominal_Amount decimal.Decimal `json:"nominal_amount" form "nominal_amount"`
	Category string `json:"category" form "category"`
	Date_Income time.Time `json:"date_income" form "date_income"`
}

type Expense struct {
	Expense_ID int `json:"expense_id" form "expense_id"`
	Kas_ID int `json:"kas_id" form "kas_id"`
	Nominal_Amount decimal.Decimal `json:"nominal_amount" form "nominal_amount"`
	Category string `json:"category" form "category"`
	Date_Expense time.Time `json:"date_expense" form "date_expense"`
}

type Kas struct {
	Kas_ID int `json:"kas_id" form "kas_id"`
	Income_ID int `json:"income_id" form "income_id"`
	Expense_ID int `json:"expense_id" form "expense_id"`
	Total int `json:"total" form "total"`
}

type Template_Savings struct {
	Template_ID int `json:"template_id" form "template_id"`
	Template_Name string `json:"template_name" form "template_name"`
	Time_Period time.Time `json:"time_period" form "time_period"`
	Additional_Information string `json:"additional_information" form "additional_information"`
}

type Savings_Progress struct {
	Savings_ID int `json:"savings_id" form "savings_id"`
	Template_ID int `json:"template_id" form "template_id"`
	ID int `json:"id" form "id"`
	Amount_Saved decimal.Decimal `json:"amount_saved" form "amount_saved"`
	Category string `json:"category" form "category"`
	Target_Date time.Time `json:"target_date" form "target_date"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{}, &Income{}, &Expense{}, &Kas{}, &Template_Savings{}, &Savings_Progress{})
}

func main() {
	db := config.InitDB()
	e := echo.New()

	//Route / to handler function
	e.Get("/users", GetUsersController)
	e.POST("/users", CreateUsersController)

	e.Start(":8000")
}

func CreateUsersController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user": user,
	})
}
	middleware.Logmiddleware(e)

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
