package router

import (
	"log"
	"net/http"
	"tictactoe/controllers"
	"tictactoe/db"
)

func StartServer() {
	playerDB := new(db.PlayerDB)
	sessionDB := new(db.SessionDB)
	gameDB := new(db.GameDB)

	loginAPI := controllers.NewLoginAPI(playerDB, sessionDB)
	gameAPI := controllers.NewGameAPI(gameDB, playerDB, sessionDB)

	router := http.NewServeMux()
	router.HandleFunc("/player/login", loginAPI.HandleLogin)
	router.HandleFunc("/player/signup", loginAPI.HandleSignup)
	router.HandleFunc("/player/token", loginAPI.HandleTokenLogin)

	router.HandleFunc("/game/new", gameAPI.HandleNewGame)
	router.HandleFunc("/game/join", gameAPI.HandleJoinGame)
	router.HandleFunc("/game/play", gameAPI.HandlePlayTurn)
	router.HandleFunc("/game/get", gameAPI.HandleGetGame)

	log.Println("started server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
