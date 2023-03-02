class GameLogic {
  static checkWinner(cells) {
    if (this.checkIfPlayerWon(cells, "X")) {
      return "X";
    } else if (this.checkIfPlayerWon(cells, "O")) {
      return "O";
    }
    return "";
  }

  static checkIfPlayerWon(cells, playerType) {
    let firstDiagonal =
      cells[0] === cells[4] && cells[4] === cells[8] && cells[0] === playerType;

    let secondDiagonal =
      cells[2] === cells[4] && cells[4] === cells[6] && cells[2] === playerType;

    let firstRow =
      cells[0] === cells[1] && cells[1] === cells[2] && cells[1] === playerType;

    let secondRow =
      cells[3] === cells[4] && cells[4] === cells[5] && cells[3] === playerType;

    let thirdRow =
      cells[6] === cells[7] && cells[7] === cells[8] && cells[6] === playerType;

    let firstColumn =
      cells[0] === cells[3] && cells[3] === cells[6] && cells[0] === playerType;

    let secondColumn =
      cells[1] === cells[4] && cells[4] === cells[7] && cells[1] === playerType;

    let thirdColumn =
      cells[2] === cells[5] && cells[5] === cells[8] && cells[8] === playerType;

    return (
      firstColumn ||
      secondColumn ||
      thirdColumn ||
      firstRow ||
      secondRow ||
      thirdRow ||
      firstDiagonal ||
      secondDiagonal
    );
  }
}

export default GameLogic;
