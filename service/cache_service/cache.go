package cache_service

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

func hash(s interface{}) (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write([]byte(b))
	return hex.EncodeToString(h.Sum(nil)), nil
}
