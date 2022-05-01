package fst_stock_service

import (
	"time"
	"yslinear/go-covid19/models"
)

type Hospital struct {
	Code     int
	Name     string
	City     string
	District string
	Address  string
	Lng      float32
	Lat      float32
	Phone    string
}

type FstStock struct {
	Hospital  Hospital
	Brand     string
	Amount    string
	Remark    string
	CreatedAt time.Time
}

func (f *FstStock) Add() error {
	hospital := map[string]interface{}{
		"code":     f.Hospital.Code,
		"name":     f.Hospital.Name,
		"city":     f.Hospital.City,
		"district": f.Hospital.District,
		"address":  f.Hospital.Address,
		"lng":      f.Hospital.Lng,
		"lat":      f.Hospital.Lat,
		"phone":    f.Hospital.Phone,
	}

	hospitalisExist, _ := models.ExistHospitalByCode(f.Hospital.Code)
	if !hospitalisExist {
		if err := models.AddHospital(hospital); err != nil {
			return err
		}
	}

	fst := map[string]interface{}{
		"hospitalCode": f.Hospital.Code,
		"brand":        f.Brand,
		"amount":       f.Amount,
		"remark":       f.Remark,
		"createdAt":    f.CreatedAt,
	}

	FstisExist, _ := models.ExistFstByHospitalCodeAndCreatedAt(f.Hospital.Code, f.CreatedAt)
	if !FstisExist {
		if err := models.AddFst(fst); err != nil {
			return err
		}
	}
	return nil
}
