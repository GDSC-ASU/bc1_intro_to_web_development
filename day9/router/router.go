package router

import (
	"net/http"
	"tictactoe/controllers"
	"tictactoe/db"
)

func StartServer() {
	playerDB := new(db.PlayerDB)

	loginAPI := controllers.NewLoginAPI(playerDB)

	router := http.NewServeMux()
	router.HandleFunc("/player/login", loginAPI.HandleLogin)
	router.HandleFunc("/player/signup", loginAPI.HandleSignup)

	http.ListenAndServe(":8080", router)
}
