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

type participantPayload struct {
	Email string `json:"email"`
}

type groupPayload struct {
	Class     string `json:"class"`
	GroupName string `json:"group_name"`
}

type votePayload struct {
	IDParticipant int    `json:"id_participant"`
	IDGroup       int    `json:"id_group"`
	Class         string `json:"class,omitempty"`
	GroupName     string `json:"group_name,omitempty"`
	Score         int    `json:"score,omitempty"`
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

func DeleteParticipant(w http.ResponseWriter, r *http.Request) {
	var payload participantPayload = participantPayload{}
	err := ReadJSON(w, r, &payload)
	if err != nil {

	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = db.DeleteParticipant(ctx, payload.Email)
	if err != nil {
		ErrorJSON(w, err, http.StatusInternalServerError)
	}

	response := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Deleted user sucessfully"),
	}
	WriteJSON(w, http.StatusOK, response)

}

func InsertGroup(w http.ResponseWriter, r *http.Request) {
	var group groupPayload = groupPayload{}

	err := ReadJSON(w, r, &group)
	if err != nil {
		ErrorJSON(w, err, http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = db.InsertGroup(ctx, group.GroupName, group.Class)
	if err != nil {
		ErrorJSON(w, err)
	}

	data := jsonResponse{
		Error:   false,
		Message: "Group Created",
	}

	WriteJSON(w, http.StatusCreated, data)

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	var group groupPayload = groupPayload{}
	err := ReadJSON(w, r, &group)
	if err != nil {
		ErrorJSON(w, err, http.StatusBadRequest)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	err = db.DeleteGroup(ctx, group.GroupName)
	if err != nil {
		ErrorJSON(w, err, http.StatusInternalServerError)
	}

	response := jsonResponse{
		Error:   false,
		Message: "Group successfully deleted",
	}

	WriteJSON(w, http.StatusOK, response)
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	var group groupPayload = groupPayload{}
	err := ReadJSON(w, r, &group)
	if err != nil {
		ErrorJSON(w, err, http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	groupData, err := db.GetGroup(ctx, group.GroupName)
	if err != nil {
		ErrorJSON(w, err, http.StatusInternalServerError)
	}

	response := jsonResponse{
		Error:   false,
		Message: "Success",
		Data:    groupData,
	}
	WriteJSON(w, http.StatusOK, response)
}

func GetGroups(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	groups, err := db.GetGroups(ctx)
	if err != nil {
		ErrorJSON(w, err, http.StatusInternalServerError)
	}

	var response = jsonResponse{
		Error:   false,
		Message: "Success",
		Data:    groups,
	}

	WriteJSON(w, http.StatusOK, response)

}

func InsertVote(w http.ResponseWriter, r *http.Request) {
	var votePayload *votePayload = &votePayload{}

	err := ReadJSON(w, r, &votePayload)
	if err != nil {
		ErrorJSON(w, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	fmt.Println(votePayload)

	err = db.InsertVote(ctx, votePayload.IDParticipant,
		votePayload.IDGroup, votePayload.Class,
		votePayload.GroupName, votePayload.Score)

	if err != nil {
		ErrorJSON(w, err)
	}

	response := jsonResponse{
		Error:   false,
		Message: "Success",
	}

	WriteJSON(w, http.StatusCreated, response)
}

func DeleteVote(w http.ResponseWriter, r *http.Request) {
	var votePayload *votePayload = &votePayload{}
	err := ReadJSON(w, r, &votePayload)
	if err != nil {

	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	err = db.DeleteVote(ctx, votePayload.IDParticipant, votePayload.IDGroup)
	if err != nil {
		ErrorJSON(w, err)
	}

	response := jsonResponse{
		Error:   false,
		Message: "Success",
	}

	WriteJSON(w, http.StatusOK, response)

}

func GetVote(w http.ResponseWriter, r *http.Request) {
	var votePayload *votePayload = &votePayload{}
	err := ReadJSON(w, r, &votePayload)
	if err != nil {
		ErrorJSON(w, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	vote, err := db.GetVoteById(ctx, votePayload.IDParticipant, votePayload.IDGroup)

	if err != nil {
		ErrorJSON(w, err)
	}

	response := jsonResponse{
		Error:   false,
		Message: "Success",
		Data:    vote,
	}

	WriteJSON(w, http.StatusOK, response)

}

func GetVotes(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	votes, err := db.GetVotes(ctx)
	if err != nil {
		ErrorJSON(w, err)
	}

	response := jsonResponse{
		Error:   false,
		Message: "Success",
		Data:    votes,
	}

	WriteJSON(w, http.StatusOK, response)
}
