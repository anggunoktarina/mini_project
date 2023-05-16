package usecase

import (
	"mini_project/model"
	"mini_project/repository/database"
)

func GetKasById(id uint) (*model.Kas, error) {
	kas, err := database.GetKasById(id)
	if err != nil {
		return nil, err
	}
	return kas, nil
}

func CreateKas(kas *model.Kas) error {
	err := database.CreateKas(kas)
	if err != nil {
		return err
	}
	return nil
}

func UpdateKas(kas *model.Kas) error {
	err := database.UpdateKas(kas)
	if err != nil {
		return err
	}
	return nil
}

// func GetRemainingBalance(kas *model.Kas) (float64, error) {
// 	total := kas.Balance
// 	incomes, err := database.GetIncomes()
// 	if err != nil {
// 		return 0, err
// 	}
// 	for _, income := range incomes {
// 		total += income.Balance
// 	}
// 	expense, err := database.GetExpense()
// 	if err != nil {
// 		return 0, err
// 	}
// 	for _, e := range expense {
// 		total -= float64(e.Balance)
// 	}
// // 	return total, nil
// }

func GetRemainingBalance(kas *model.Kas) (float64, error) {
	total := kas.Balance
	income, err := database.GetIncomes()
	if err != nil {
		return 0, err
	}
	for _, i := range income {
		total += i.Balance
	}
	expense, err := database.GetExpense()
	if err != nil {
		return 0, err
	}
	for _, e := range expense {
		total -= e.Balance
	}
	return float64(total), nil
}
