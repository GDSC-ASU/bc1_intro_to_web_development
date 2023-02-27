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

	loginAPI := controllers.NewLoginAPI(playerDB, sessionDB)

	router := http.NewServeMux()
	router.HandleFunc("/player/login", loginAPI.HandleLogin)
	router.HandleFunc("/player/signup", loginAPI.HandleSignup)
	router.HandleFunc("/player/token", loginAPI.HandleTokenLogin)

	log.Println("started server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
