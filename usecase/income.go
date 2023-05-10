package usecase

import (
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

func CreateIncome(income *model.Income) error {

	err := database.CreateIncome(income)
	if err != nil {
		return err
	}

	return nil
}

func GetIncome(id uint) (income model.Income, err error) {
	income, err = database.GetIncome(id)
	if err != nil {
		fmt.Println("GetIncome: Error getting income from database")
		return
	}
	return
}

func GetListIncomes() (incomes []model.Income, err error) {
	incomes, err = database.GetIncomes()
	if err != nil {
		fmt.Println("GetListIncomes: Error getting incomes from database")
		return
	}
	return
}

func UpdateIncome(income *model.Income) (err error) {
	err = database.UpdateIncome(income)
	if err != nil {
		fmt.Println("UpdateIncome : Error updating income, err: ", err)
		return
	}

	return
}

func DeleteIncome(id uint) (err error) {
	income := model.Income{}
	income.ID = id
	err = database.DeleteIncome(&income)
	if err != nil {
		fmt.Println("DeleteIncome : error deleting income, err: ", err)
		return
	}

	return
}
