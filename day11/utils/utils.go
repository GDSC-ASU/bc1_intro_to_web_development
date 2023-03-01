package utils

func CheckWinner(cells []string) string {
	if checkIfPlayerWon(cells, "X") {
		return "X"
	} else if checkIfPlayerWon(cells, "O") {
		return "O"
	}
	return ""
}

func checkIfPlayerWon(cells []string, playerType string) bool {
	var firstDiagonal = cells[0] == cells[4] && cells[4] == cells[8] && cells[0] == playerType

	var secondDiagonal = cells[2] == cells[4] && cells[4] == cells[6] && cells[2] == playerType

	var firstRow = cells[0] == cells[1] && cells[1] == cells[2] && cells[1] == playerType

	var secondRow = cells[3] == cells[4] && cells[4] == cells[5] && cells[3] == playerType

	var thirdRow = cells[6] == cells[7] && cells[7] == cells[8] && cells[6] == playerType

	var firstColumn = cells[0] == cells[3] && cells[3] == cells[6] && cells[0] == playerType

	var secondColumn = cells[1] == cells[4] && cells[4] == cells[7] && cells[1] == playerType

	var thirdColumn = cells[2] == cells[5] && cells[5] == cells[8] && cells[8] == playerType

	return firstColumn ||
		secondColumn ||
		thirdColumn ||
		firstRow ||
		secondRow ||
		thirdRow ||
		firstDiagonal ||
		secondDiagonal
}
