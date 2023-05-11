package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreateKas(kas *model.Kas) error {
	if err := config.DB.Create(kas).Error; err != nil {
		return err
	}
	return nil
}

//	func GetKases() (kases []model.Kas, err error) {
//		if err = config.DB.Find(&kases).Error; err != nil {
//			return
//		}
//		return
//	}
func GetKasById(id uint) (*model.Kas, error) {
	kas := &model.Kas{}
	if err := config.DB.First(kas, id).Error; err != nil {
		return nil, err
	}
	return kas, nil
}

func UpdateKas(kas *model.Kas) error {
	if err := config.DB.Updates(kas).Error; err != nil {
		return err
	}
	return nil
}

func DeleteKas(kas *model.Kas) error {
	if err := config.DB.Delete(kas).Error; err != nil {
		return err
	}
	return nil
}
