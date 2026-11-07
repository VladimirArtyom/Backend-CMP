package internal

import (
	"compro-exh/internal/db"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 104876
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	var dec *json.Decoder = json.NewDecoder(r.Body)

	err := dec.Decode(data)

	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("Body must have only a signle JSON value")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	outByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(outByte)
	if err != nil {
		return err
	}

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	var statusCode int = http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return WriteJSON(w, statusCode, payload)

}

type participantPayload struct {
	Email string `json:"email"`
}

func InsertParticipant(w http.ResponseWriter, r *http.Request) {

	var payload participantPayload = participantPayload{}
	err := ReadJSON(w, r, &payload)

	if err != nil {
		ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.InsertParticipant(ctx, payload.Email)
	if err != nil {
		ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("User inserted: %v", payload.Email),
	}
	WriteJSON(w, http.StatusCreated, response)
}

func GetParticipant(w http.ResponseWriter, r *http.Request) {

	var payload *participantPayload = &participantPayload{}
	err := ReadJSON(w, r, &payload)
	if err != nil {
		ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	participant, err := db.GetParticipant(ctx, payload.Email)
	if err != nil {
		ErrorJSON(w, err, http.StatusInternalServerError)
	}

	response := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Success"),
		Data:    participant,
	}

	WriteJSON(w, http.StatusOK, response)

}
