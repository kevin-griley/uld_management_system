package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kevin-griley/api/data"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandlePostUser(w http.ResponseWriter, r *http.Request) *ApiError {
	registerReq := new(LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(registerReq); err != nil {
		return &ApiError{http.StatusBadRequest, err.Error()}
	}

	user, err := data.NewUser(registerReq.Email, registerReq.Password)
	if err != nil {
		return &ApiError{http.StatusBadRequest, err.Error()}
	}

	userID, err := data.CreateUser(user)
	if err != nil {
		return &ApiError{http.StatusInternalServerError, err.Error()}
	}

	user.ID = userID
	return WriteJSON(w, http.StatusOK, user)

}
