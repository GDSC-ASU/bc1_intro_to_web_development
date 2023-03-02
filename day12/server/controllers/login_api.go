package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tictactoe/db"
	"tictactoe/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginAPI struct {
	playerDB  *db.PlayerDB
	sessionDB *db.SessionDB
}

func NewLoginAPI(playerDB *db.PlayerDB, sessionDB *db.SessionDB) *LoginAPI {
	return &LoginAPI{
		playerDB:  playerDB,
		sessionDB: sessionDB,
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

	player.Password = ""

	json.NewEncoder(resp).Encode(player)
}

func (l *LoginAPI) HandleLogin(resp http.ResponseWriter, req *http.Request) {
	var requestBody reqBody

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(resp, "👎👎 your request has an incorrect form 👎👎")
		return
	}

	player, err := l.playerDB.Get(requestBody.Username)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(resp, "username is incorrect 👎")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(requestBody.Password))
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(resp, "password is incorrect 👎")
		return
	}

	session := models.Session{
		Username: player.Username,
		Token:    strings.ReplaceAll(uuid.NewString(), "-", ""),
	}

	err = l.sessionDB.Create(session)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(resp, "something went wrong ☹️")
		return
	}

	json.NewEncoder(resp).Encode(map[string]any{
		"token": session.Token,
	})
}

func (l *LoginAPI) HandleTokenLogin(resp http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Authorization")

	session, err := l.sessionDB.Get(token)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(resp, "invalid token 👎")
		return
	}

	player, err := l.playerDB.Get(session.Username)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(resp, "something went wrong")
		return
	}

	player.Password = ""

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(player)
}
