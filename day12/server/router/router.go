package router

import (
	"log"
	"net/http"
	"tictactoe/controllers"
	"tictactoe/db"
)

func ignoreOptionsMethod(handler http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Authorization")
		resp.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		resp.Header().Set("Content-Type", "application/json")

		if req.Method != http.MethodOptions {
			handler.ServeHTTP(resp, req)
		}
	}
}

func StartServer() {
	playerDB := new(db.PlayerDB)
	sessionDB := new(db.SessionDB)
	gameDB := new(db.GameDB)

	loginAPI := controllers.NewLoginAPI(playerDB, sessionDB)
	gameAPI := controllers.NewGameAPI(gameDB, playerDB, sessionDB)

	router := http.NewServeMux()
	router.HandleFunc("/player/login", ignoreOptionsMethod(loginAPI.HandleLogin))
	router.HandleFunc("/player/signup", ignoreOptionsMethod(loginAPI.HandleSignup))
	router.HandleFunc("/player/token", ignoreOptionsMethod(loginAPI.HandleTokenLogin))

	router.HandleFunc("/game/new", ignoreOptionsMethod(gameAPI.HandleNewGame))
	router.HandleFunc("/game/join", ignoreOptionsMethod(gameAPI.HandleJoinGame))
	router.HandleFunc("/game/play", ignoreOptionsMethod(gameAPI.HandlePlayTurn))
	router.HandleFunc("/game/get", ignoreOptionsMethod(gameAPI.HandleGetGame))

	log.Println("started server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
