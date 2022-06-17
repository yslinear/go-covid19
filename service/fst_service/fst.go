package fst_service

import (
	"encoding/json"
	"time"
	"yslinear/go-covid19/models"
	"yslinear/go-covid19/pkg/gredis"
	"yslinear/go-covid19/service/cache_service"
)

type Fst struct {
	HospitalCode string
	Brand        string
	Amount       int
	Remark       string
	CreatedAt    time.Time

	PageNum  int
	PageSize int
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
	var cacheFsts []*models.Fst

	cache := cache_service.Fst(*f)
	key, err := cache.GetFstsKey()
	if err != nil {
		return nil, err
	}
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			return nil, err
		} else {
			json.Unmarshal(data, &cacheFsts)
			return cacheFsts, nil
		}
	}

	fsts, err := models.GetFsts(f.PageNum, f.PageSize, f.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, fsts, 3600)
	return fsts, nil
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
