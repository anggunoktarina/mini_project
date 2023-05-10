package usecase

import (
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

func CreateExpense(expense *model.Expense) error {

	err := database.CreateExpense(expense)
	if err != nil {
		return err
	}

	return nil
}

func GetExpense(expense *model.Expense) error {
	err := database.GetExpense(expense)
	if err != nil {
		fmt.Println("GetExpense: Error getting expense from database")
		return err
	}
	return nil
}

func UpdateExpense(expense *model.Expense) (err error) {
	err = database.UpdateExpense(expense)
	if err != nil {
		fmt.Println("UpdateExpense : Error updating expense, err: ", err)
		return
	}

	return
}

func DeleteExpense(id uint) (err error) {
	expense := model.Expense{}
	expense.ID = id
	err = database.DeleteExpense(&expense)
	if err != nil {
		fmt.Println("DeleteExpense : error deleting expense, err: ", err)
		return
	}

	return
}
