package models

import "gorm.io/gorm"

type Hospital struct {
	gorm.Model
	Code    int
	Name    string
	Address string
	Lng     float32
	Lat     float32
	Phone   string
}

func AddHospital(data map[string]interface{}) error {
	hospital := Hospital{
		Code:    data["code"].(int),
		Name:    data["name"].(string),
		Address: data["address"].(string),
		Lng:     data["lng"].(float32),
		Lat:     data["lat"].(float32),
		Phone:   data["phone"].(string),
	}
	if err := db.Create(&hospital).Error; err != nil {
		return err
	}

	return nil
}

func ExistHospitalByCode(code int) (bool, error) {
	var hospital Hospital

	if err := db.Select("id").Where("code = ?", code).First(&hospital).Error; err != nil {
		return false, err
	}
	if hospital.ID > 0 {
		return true, nil
	}

	return false, nil
}
