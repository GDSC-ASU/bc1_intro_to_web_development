package main

import (
	"tictactoe/db"
)

func main() {
	playerDB := new(db.PlayerDB)
	err := playerDB.Delete("Laith")
	if err != nil {
		panic(err)
	}
}
