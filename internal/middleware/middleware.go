package middleware

import (
	models "L0/internal"
	"encoding/json"
	"errors"
)

func CheckModel(body []byte) (string, error) {
	var model models.Model
	err := json.Unmarshal(body, &model)
	if err != nil {
		return "", err
	}

	if model.OrderUid == "" || model.TrackNumber == "" || model.Payment.Transaction == "" {
		err := errors.New("uncorrected model")
		return "", err

	}
	return model.OrderUid, nil
}
