package controllers

import (
	"encoding/json"
	"net/http"
	"tictactoe/db"
	"tictactoe/models"

	"golang.org/x/crypto/bcrypt"
)

type LoginAPI struct {
	playerDB *db.PlayerDB
}

func NewLoginAPI(playerDB *db.PlayerDB) *LoginAPI {
	return &LoginAPI{
		playerDB: playerDB,
	}
}

type reqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginAPI) HandleSignup(resp http.ResponseWriter, req *http.Request) {
	var requestBody reqBody

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.MinCost)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	requestBody.Password = string(hashed)

	player, err := l.playerDB.Create(models.Player{
		Username: requestBody.Username,
		Password: requestBody.Password,
	})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(player)
}

func (l *LoginAPI) HandleLogin(resp http.ResponseWriter, req *http.Request) {

}
