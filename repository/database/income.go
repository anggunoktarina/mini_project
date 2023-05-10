package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreateIncome(income *model.Income) error {
	if err := config.DB.Create(income).Error; err != nil {
		return err
	}
	return nil
}

func GetIncomes() (incomes []model.Income, err error) {
	if err = config.DB.Find(&incomes).Error; err != nil {
		return
	}
	return
}

func GetIncome(id uint) (income model.Income, err error) {
	income.ID = id
	if err = config.DB.First(&income).Error; err != nil {
		return
	}
	return
}

func UpdateIncome(income *model.Income) error {
	if err := config.DB.Updates(income).Error; err != nil {
		return err
	}
	return nil
}

func DeleteIncome(income *model.Income) error {
	if err := config.DB.Delete(income).Error; err != nil {
		return err
	}
	return nil
}
