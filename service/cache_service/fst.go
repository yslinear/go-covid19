package cache_service

import "time"

type Fst struct {
	HospitalCode string
	Brand        string
	Amount       int
	Remark       string
	CreatedAt    time.Time

	PageNum  int
	PageSize int
}

func (f *Fst) GetFstsKey() (string, error) {
	hash, err := hash(f)
	if err != nil {
		return "", err
	}
	return "FSTS:" + hash, nil
}
