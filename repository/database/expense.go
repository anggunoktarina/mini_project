package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreateExpense(expense *model.Expense) error {
	if err := config.DB.Create(expense).Error; err != nil {
		return err
	}
	return nil
}

func GetExpense() (expense []model.Expense, err error) {
	if err = config.DB.Find(&expense).Error; err != nil {
		return
	}
	return
}

func UpdateExpense(expense *model.Expense) error {
	if err := config.DB.Updates(expense).Error; err != nil {
		return err
	}
	return nil
}

func DeleteExpense(expense *model.Expense) error {
	if err := config.DB.Delete(expense).Error; err != nil {
		return err
	}
	return nil
}
