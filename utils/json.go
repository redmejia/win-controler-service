package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) error {

	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(dataByte)
	if err != nil {
		return err
	}

	return nil
}

func ReadJSON(r *http.Request, data any) error {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(data)
	if err != nil {
		return err
	}

	return nil
}
