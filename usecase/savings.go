package usecase

import (
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

func CreateSavings(savings *model.Savings) error {

	err := database.CreateSavings(savings)
	if err != nil {
		return err
	}

	return nil
}

func GetAllSavings() ([]model.Savings, error) {
	savings, err := database.GetSavings()
	if err != nil {
		return nil, err
	}
	return savings, nil
}

// func GetSavings(savings *model.Savings) error {
// 	err := database.GetSavings(savings)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func GetSavingsByID(id uint) ([]model.Savings, error) {
	savings, err := database.GetSavings()
	if err != nil {
		fmt.Println("GetSavingsByID: Error getting savings from database")
		return nil, err
	}

	var result []model.Savings
	for _, s := range savings {
		if s.ID == id {
			result = append(result, s)
		}
	}

	return result, nil
}

// func GetSavingsByID(id uint) (savings model.Savings, err error) {
// 	savings, err = database.GetSavings(id)
// 	if err != nil {
// 		fmt.Println("GetSavings: Error getting savings from database")
// 		return model.Savings{}, err
// 	}
// 	return savings, nil
// }

func UpdateSavings(savings *model.Savings) (err error) {
	err = database.UpdateSavings(savings)
	if err != nil {
		fmt.Println("UpdateSavings: Error updating savings, err: ", err)
		return
	}

	return
}

func DeleteSavings(id uint) (err error) {
	savings := model.Savings{}
	savings.ID = id
	err = database.DeleteSavings(&savings)
	if err != nil {
		fmt.Println("DeleteSavings : error deleting savings, err: ", err)
		return
	}

	return
}
