package hospital_service

import "yslinear/go-covid19/models"

type Zone struct {
	City     string
	District string
}

func GetAllCities() ([]string, error) {
	return models.GetAllHospitalCities()
}

func GetAllDistricts(city string) ([]string, error) {
	return models.GetAllHospitalDistricts(city)
}
