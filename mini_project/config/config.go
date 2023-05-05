package config

import (
	"fmt"
	"mini_project_anggun/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// "DB_Username": "root",
// "DB_Password": "Anggun20",
// "DB_Port":     "3306",
// "DB_Host":     "altadb.c3fzsopw9aux.ap-southeast-1.rds.amazonaws.com",
// "DB_Name":     "structure_project",
func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "Anggun20",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "mini_project",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Income{}, &model.Expense{}, &model.Kas{}, &model.Template_Savings{}, &model.Savings_Progress{})
}
