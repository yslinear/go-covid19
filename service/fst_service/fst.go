package fst_service

import (
	"time"
	"yslinear/go-covid19/models"
)

type Fst struct {
	HospitalCode string
	Brand        string
	Amount       int
	Remark       string
	CreatedAt    time.Time
}

func (f *Fst) Add() error {
	fst := map[string]interface{}{
		"hospitalCode": f.HospitalCode,
		"brand":        f.Brand,
		"amount":       f.Amount,
		"remark":       f.Remark,
		"createdAt":    f.CreatedAt,
	}

	if err := models.AddFst(fst); err != nil {
		return err
	}

	return nil
}

func (f *Fst) Count() (int64, error) {
	return models.GetFstTotal(f.getMaps())
}

func (f *Fst) GetAll() ([]*models.Fst, error) {
	fsts, err := models.GetFsts(f.getMaps())
	if err != nil {
		return nil, err
	}
	return fsts, nil
}

func (f *Fst) Get() ([]*models.Fst, error) {
	fst, err := models.GetFst(f.getMaps())
	if err != nil {
		return nil, err
	}
	return fst, nil
}

func (f *Fst) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	if f.HospitalCode != "" {
		maps["hospital_code"] = f.HospitalCode
	}
	if !f.CreatedAt.IsZero() {
		maps["created_at"] = f.CreatedAt
	}

	return maps
}
