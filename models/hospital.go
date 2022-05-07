package models

import "gorm.io/gorm"

type Hospital struct {
	gorm.Model
	Code     int
	Name     string
	City     string
	District string
	Address  string
	Lng      float64
	Lat      float64
	Phone    string
}

func AddHospital(data map[string]interface{}) error {
	hospital := Hospital{
		Code:     data["code"].(int),
		Name:     data["name"].(string),
		City:     data["city"].(string),
		District: data["district"].(string),
		Address:  data["address"].(string),
		Lng:      data["lng"].(float64),
		Lat:      data["lat"].(float64),
		Phone:    data["phone"].(string),
	}
	if err := db.Create(&hospital).Error; err != nil {
		return err
	}

	return nil
}

func GetHospitals(maps interface{}) ([]*Hospital, error) {
	var hospitals []*Hospital
	if err := db.Where(maps).Find(&hospitals).Error; err != nil {
		return nil, err
	}

	return hospitals, nil
}

func GetHospitalTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Hospital{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
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

func GetAllHospitalCities() ([]string, error) {
	var cities []string
	if err := db.Table("hospitals").Select("city").Group("city").Pluck("city", &cities).Error; err != nil {
		return nil, err
	}

	return cities, nil
}

func GetAllHospitalDistricts(city string) ([]string, error) {
	var districts []string
	if err := db.Table("hospitals").Select("district").Where("city = ?", city).Group("district").Pluck("district", &districts).Error; err != nil {
		return nil, err
	}

	return districts, nil
}

func GetHospital(maps interface{}) (*Hospital, error) {
	var hospital *Hospital
	if err := db.Where(maps).Find(&hospital).Error; err != nil {
		return nil, err
	}

	return hospital, nil
}
