package models

import (
	"errors"
	"fmt"
	"strings"
	"tictactoe/utils"
)

type Game struct {
	Id              int    `json:"id"`
	XPlayerId       int    `json:"xPlayerId"`
	OPlayerId       int    `json:"oPlayerId"`
	XScore          int    `json:"xScore"`
	OScore          int    `json:"oScore"`
	Cells           string `json:"cells"`
	CurrentPlayerId int    `json:"currentPlayerId"`
}

func (g *Game) resetGame() {
	g.Cells = "1 2 3 4 5 6 7 8 9"
	g.CurrentPlayerId = g.XPlayerId
}

func (g *Game) switchTurns() {
	if g.CurrentPlayerId == g.XPlayerId {
		g.CurrentPlayerId = g.OPlayerId
	} else if g.CurrentPlayerId == g.OPlayerId {
		g.CurrentPlayerId = g.XPlayerId
	}
}

func (g *Game) MarkCell(cellNumber int) error {
	if !g.isCellNumberInBounds(cellNumber) {
		return errors.New("cell number is not in bound")
	}

	cells := strings.Split(g.Cells, " ")
	if g.areAllCellsOccupied(cells) {
		return errors.New("all game's cells are occupied")
	}

	if g.isCellOccupied(cells, cellNumber) {
		return errors.New("the requested cell: " + fmt.Sprint(cellNumber) + " is occupied with: " + cells[cellNumber])
	}

	currentPlayer := g.getCurrentPlayerSymbol()
	cells[cellNumber] = currentPlayer
	g.switchTurns()

	return nil
}

func (g *Game) CheckWinnerAndUpdateScore() {
	cells := strings.Split(g.Cells, " ")

	if g.areAllCellsOccupied(cells) {
		g.resetGame()
		return
	}

	winner := utils.CheckWinner(cells)
	switch winner {
	case "X":
		g.XScore++
	case "O":
		g.OScore++
	}

	if len(winner) > 0 {
		g.resetGame()
	}
}

func (g *Game) isCellOccupied(cells []string, cellNumber int) bool {
	return cells[cellNumber] != "X" && cells[cellNumber] != "O"
}

func (g *Game) areAllCellsOccupied(cells []string) bool {
	for index, _ := range cells {
		if !g.isCellOccupied(cells, index) {
			return false
		}
	}
	return true
}

func (g *Game) isCellNumberInBounds(cellNumber int) bool {
	return cellNumber < 9 && cellNumber >= 0
}

func (g *Game) getCurrentPlayerSymbol() string {
	if g.CurrentPlayerId == g.XPlayerId {
		return "X"
	} else if g.CurrentPlayerId == g.OPlayerId {
		return "O"
	}
	return ""
}
