package models

import (
	"time"

	"gorm.io/gorm"
)

type Fst struct {
	gorm.Model
	HospitalCode int
	Brand        string
	Amount       int
	Remark       string
	CreatedAt    time.Time
}

func AddFst(data map[string]interface{}) error {
	fst := Fst{
		HospitalCode: data["hospitalCode"].(int),
		Brand:        data["brand"].(string),
		Amount:       data["amount"].(int),
		Remark:       data["remark"].(string),
		CreatedAt:    data["createdAt"].(time.Time),
	}
	if err := db.Create(&fst).Error; err != nil {
		return err
	}

	return nil
}

func GetFsts(maps interface{}) ([]*Fst, error) {
	var fsts []*Fst
	if err := db.Where(maps).Find(&fsts).Error; err != nil {
		return nil, err
	}

	return fsts, nil
}

func GetFstTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Fst{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func ExistFstByHospitalCodeAndCreatedAt(hospitalCode int, createdAt time.Time) (bool, error) {
	var fst Fst
	if err := db.Where("hospital_code = ? AND created_at = ?", hospitalCode, createdAt).First(&fst).Error; err != nil {
		return false, err
	}
	if fst.ID > 0 {
		return true, nil
	}

	return false, nil
}
