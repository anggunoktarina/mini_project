package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreateSavings(savings *model.Savings) error {
	if err := config.DB.Create(savings).Error; err != nil {
		return err
	}
	return nil
}

func GetSavings() (savings []model.Savings, err error) {
	if err = config.DB.Find(&savings).Error; err != nil {
		return
	}
	return
}

func GetSavingsByID(id uint) (savings model.Savings, err error) {
	savings.ID = id
	if err = config.DB.First(savings).Error; err != nil {
		return
	}
	return
}

func UpdateSavings(savings *model.Savings) error {
	if err := config.DB.Updates(savings).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSavings(savings *model.Savings) error {
	if err := config.DB.Delete(savings).Error; err != nil {
		return err
	}
	return nil
}
