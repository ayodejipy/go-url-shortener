package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ParseJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return errors.New("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

// Error Messages
func WriteMessage(w http.ResponseWriter, code int, payload Response) error {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	return nil
}

func SuccessMessage(w http.ResponseWriter, statusCode int, payload Response) {
	WriteMessage(w, statusCode, Response{
		Success: true,
		Message: payload.Message,
		Data:    payload.Data,
	})
}

func ErrorMessage(w http.ResponseWriter, code int, err error) {
	WriteMessage(w, code, Response{
		Success: false,
		Message: err.Error(),
	})
}

func InternalError(w http.ResponseWriter, err error) {
	WriteMessage(w, http.StatusInternalServerError, Response{
		Success: false,
		Message: err.Error(),
	})
}

func BadRequestError(w http.ResponseWriter, err error) {
	WriteMessage(w, http.StatusBadRequest, Response{
		Success: false,
		Message: err.Error(),
	})
}
