package hospital_service

import (
	"encoding/json"
	"yslinear/go-covid19/models"
	"yslinear/go-covid19/pkg/gredis"

	"yslinear/go-covid19/service/cache_service"
)

type Hospital struct {
	Code     string
	Name     string
	City     string
	District string
	Address  string
	Lng      float64
	Lat      float64
	Phone    string
}

func (h *Hospital) Add() error {
	hospital := map[string]interface{}{
		"code":     h.Code,
		"name":     h.Name,
		"city":     h.City,
		"district": h.District,
		"address":  h.Address,
		"lng":      h.Lng,
		"lat":      h.Lat,
		"phone":    h.Phone,
	}

	if err := models.AddHospital(hospital); err != nil {
		return err
	}

	return nil
}

func (h *Hospital) Count() (int64, error) {
	return models.GetHospitalTotal(h.getMaps())
}

func (h *Hospital) GetAll() ([]*models.Hospital, error) {
	var cacheHospitals []*models.Hospital

	cache := cache_service.Hospital(*h)
	key, err := cache.GetHospitalsKey()
	if err != nil {
		return nil, err
	}
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			return nil, err
		} else {
			json.Unmarshal(data, &cacheHospitals)
			return cacheHospitals, nil
		}
	}

	hospitals, err := models.GetHospitals(h.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, hospitals, 3600)
	return hospitals, nil
}

func (h *Hospital) Get() (*models.Hospital, error) {
	var cacheHospital *models.Hospital

	cache := cache_service.Hospital(*h)
	key, err := cache.GetHospitalKey()
	if err != nil {
		return nil, err
	}
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			return nil, err
		} else {
			json.Unmarshal(data, &cacheHospital)
			return cacheHospital, nil
		}
	}

	hospital, err := models.GetHospital(h.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, hospital, 3600)
	return hospital, nil
}

func (h *Hospital) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	if h.Code != "" {
		maps["code"] = h.Code
	}
	if h.City != "" {
		maps["city"] = h.City
	}
	if h.District != "" {
		maps["district"] = h.District
	}

	return maps
}
