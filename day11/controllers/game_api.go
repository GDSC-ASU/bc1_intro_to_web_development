package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tictactoe/db"
	"tictactoe/models"
)

type GameAPI struct {
	gameDB    *db.GameDB
	playerDB  *db.PlayerDB
	sessionDB *db.SessionDB
}

func NewGameAPI(gameDB *db.GameDB, playerDB *db.PlayerDB, sessionDB *db.SessionDB) *GameAPI {
	return &GameAPI{
		gameDB:    gameDB,
		playerDB:  playerDB,
		sessionDB: sessionDB,
	}
}

func (g *GameAPI) setupResponseHeaders(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
}

func (g *GameAPI) HandleNewGame(resp http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Authorization")
	if len(token) == 0 {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	player, err := g.fetchPlayerUsingToken(token)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	game := models.Game{
		XPlayerId: player.Id,
	}

	newGame, err := g.gameDB.Create(game)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	g.setupResponseHeaders(resp)
	json.NewEncoder(resp).Encode(newGame)
}

func (g *GameAPI) HandleJoinGame(resp http.ResponseWriter, req *http.Request) {
	gameId := req.URL.Query().Get("id")
	if len(gameId) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	intGameId, err := strconv.Atoi(gameId)
	if len(gameId) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	token := req.Header.Get("Authorization")
	if len(token) == 0 {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	player, err := g.fetchPlayerUsingToken(token)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	game, err := g.gameDB.Get(intGameId)
	if err != nil {
		fmt.Println(err)
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	if game.OPlayerId != 0 {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(resp, "THIS GAME IS FULLY OCCUPIED!!!")
		return
	}

	game.OPlayerId = player.Id

	err = g.gameDB.Update(game.Id, game)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	g.setupResponseHeaders(resp)
	json.NewEncoder(resp).Encode(game)
}

func (g *GameAPI) HandlePlayTurn(resp http.ResponseWriter, req *http.Request) {
	var reqBody struct {
		GameId     int `json:"gameId"`
		CellNumber int `json:"cellNumber"`
	}

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	token := req.Header.Get("Authorization")
	if len(token) == 0 {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	player, err := g.fetchPlayerUsingToken(token)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	game, err := g.gameDB.Get(reqBody.GameId)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if player.Id != game.CurrentPlayerId {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(resp, "this player can't play right now")
		return
	}

	err = game.MarkCell(reqBody.CellNumber)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(resp, err.Error())
		return
	}

	game.CheckWinnerAndUpdateScore()

	err = g.gameDB.Update(game.Id, game)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	g.setupResponseHeaders(resp)
	json.NewEncoder(resp).Encode(game)
}

func (g *GameAPI) HandleGetGame(resp http.ResponseWriter, req *http.Request) {
	gameId := req.URL.Query().Get("id")
	if len(gameId) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	intGameId, err := strconv.Atoi(gameId)
	if len(gameId) == 0 {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	token := req.Header.Get("Authorization")
	if len(token) == 0 {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	player, err := g.fetchPlayerUsingToken(token)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	game, err := g.gameDB.Get(intGameId)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if player.Id != game.XPlayerId && player.Id != game.OPlayerId {
		resp.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(resp, "this player does not belong to this game!")
		return
	}

	g.setupResponseHeaders(resp)
	json.NewEncoder(resp).Encode(game)
}

func (g *GameAPI) fetchPlayerUsingToken(token string) (models.Player, error) {
	session, err := g.sessionDB.Get(token)
	if err != nil {
		return models.Player{}, err
	}

	player, err := g.playerDB.Get(session.Username)
	if err != nil {
		return models.Player{}, err
	}

	return player, nil
}
