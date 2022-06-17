package cache_service

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

func (h *Hospital) GetHospitalKey() (string, error) {
	hash, err := hash(h)
	if err != nil {
		return "", err
	}
	return "HOSPITAL:" + hash, nil
}

func (h *Hospital) GetHospitalsKey() (string, error) {
	hash, err := hash(h)
	if err != nil {
		return "", err
	}
	return "HOSPITALS:" + hash, nil
}
