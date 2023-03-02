import GameLogic from "./GameLogic";

class GameData {
  cells;
  turn;
  xScore;
  oScore;

  constructor(copiedData) {
    if (copiedData !== undefined && copiedData !== null) {
      this.cells = copiedData.cells;
      this.turn = copiedData.turn;
      this.xScore = copiedData.xScore;
      this.oScore = copiedData.oScore;
      return;
    }
    this.cells = ["1", "2", "3", "4", "5", "6", "7", "8", "9"];
    this.turn = "X";
    this.xScore = 0;
    this.oScore = 0;
  }

  resetGame() {
    this.cells = ["1", "2", "3", "4", "5", "6", "7", "8", "9"];
    this.turn = "X";
  }

  markCell(cellNumber) {
    if (this.isCellNumberInBounds(cellNumber)) return;

    if (this.isCellOccupied(cellNumber)) return;

    switch (this.turn) {
      case "X":
        this.cells[cellNumber] = "X";
        break;
      case "O":
        this.cells[cellNumber] = "O";
        break;
    }

    if (this.areAllCellsOccupied()) {
      this.resetGame();
      return;
    }

    this.switchTurn();
  }

  isCellNumberInBounds(cellNumber) {
    return cellNumber < 0 || cellNumber > this.cells.length - 1;
  }

  isCellOccupied(cellNumber) {
    return this.cells[cellNumber] === "X" || this.cells[cellNumber] === "O";
  }

  areAllCellsOccupied() {
    for (let i = 0; i < this.cells.length; i++) {
      if (!this.isCellOccupied(i)) {
        return false;
      }
    }
    return true;
  }

  switchTurn() {
    switch (this.turn) {
      case "X":
        this.turn = "O";
        break;
      case "O":
        this.turn = "X";
        break;
    }
  }

  checkWinner() {
    let winner = GameLogic.checkWinner(this.cells);
    switch (winner) {
      case "X":
        this.xScore++;
        this.resetGame();
        return "X";
      case "O":
        this.oScore++;
        this.resetGame();
        return "O";
    }
  }
}

export default GameData;
