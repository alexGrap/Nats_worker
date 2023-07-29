package middleware

import (
	models "L0/internal"
	"encoding/json"
)

func CheckModel(body []byte) (string, error) {
	var model models.Model
	err := json.Unmarshal(body, &model)
	if err != nil {
		return "", err
	}

	if model.OrderUid == "" || model.TrackNumber == "" || model.Payment.Transaction == "" {
		return "", err
	}
	return model.OrderUid, nil
}
