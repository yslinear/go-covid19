package hospital_service

import "yslinear/go-covid19/models"

type Hospital struct {
	Code     int
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
	hospitals, err := models.GetHospitals(h.getMaps())
	if err != nil {
		return nil, err
	}
	return hospitals, nil
}

func (h *Hospital) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	if h.Code != 0 {
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
