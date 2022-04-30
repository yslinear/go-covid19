package models

import "gorm.io/gorm"

type Fst struct {
	gorm.Model
	Brand  string
	Amount string
	Remark string
}

func AddFst(data map[string]interface{}) error {
	fst := Fst{
		Brand:  data["brand"].(string),
		Amount: data["amount"].(string),
		Remark: data["remark"].(string),
	}
	if err := db.Create(&fst).Error; err != nil {
		return err
	}

	return nil
}
