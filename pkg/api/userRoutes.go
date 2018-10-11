package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := CreateUserRequest{}
	err := decoder.Decode(&data)
	if err != nil {
		RespondError(w, errors.New("BAD_REQUEST"))
		return
	}

	token, err := CreateUser(data.Username, data.Password)
	if err != nil {
		RespondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(CreateUserResponse{Token: token})
}

type AuthenticateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticateUserResponse struct {
	Token string `json:"token"`
}

func AuthenticateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	data := AuthenticateUserRequest{}
	err := decoder.Decode(&data)
	if err != nil {
		RespondError(w, errors.New("BAD_REQUEST"))
		return
	}

	user, err := GetUserByUsername(data.Username)
	if err != nil {
		RespondError(w, err)
		return
	}

	correctPassword := user.checkPassword(data.Password)
	if correctPassword != true {
		RespondError(w, errors.New("Incorrect password"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(AuthenticateUserResponse{Token: user.Token})
}
