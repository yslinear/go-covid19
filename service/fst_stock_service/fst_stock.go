package fst_stock_service

import (
	"yslinear/go-covid19/models"
)

type Hospital struct {
	Code    int
	Name    string
	Address string
	Lng     float32
	Lat     float32
	Phone   string
}

type FstStock struct {
	Hospital Hospital
	Brand    string
	Amount   string
	Remark   string
}

func (f *FstStock) Add() error {
	hospital := map[string]interface{}{
		"code":    f.Hospital.Code,
		"name":    f.Hospital.Name,
		"address": f.Hospital.Address,
		"lng":     f.Hospital.Lng,
		"lat":     f.Hospital.Lat,
		"phone":   f.Hospital.Phone,
	}

	fst := map[string]interface{}{
		"brand":  f.Brand,
		"amount": f.Amount,
		"remark": f.Remark,
	}

	if err := models.AddHospital(hospital); err != nil {
		return err
	}
	if err := models.AddFst(fst); err != nil {
		return err
	}

	return nil
}
