package hospital_service

import "yslinear/go-covid19/models"

type Hospital struct {
	Code    int
	Name    string
	Address string
	Lng     float32
	Lat     float32
	Phone   string
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
	return maps
}
