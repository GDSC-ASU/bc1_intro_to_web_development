class Game {
  id;
  xPlayerId;
  oPlayerId;
  xScore;
  oScore;
  cells;
  currentPlayerId;

  constructor(copiedGame) {
    if (copiedGame) {
      this.id = copiedGame.id;
      this.xPlayerId = copiedGame.xPlayerId;
      this.oPlayerId = copiedGame.oPlayerId;
      this.xScore = copiedGame.xScore;
      this.oScore = copiedGame.oScore;
      this.cells = copiedGame.cells.split(" ");
      this.currentPlayerId = copiedGame.currentPlayerId;
      return;
    }

    this.id = 0;
    this.xPlayerId = 0;
    this.oPlayerId = 0;
    this.xScore = 0;
    this.oScore = 0;
    this.cells = ["1", "2", "3", "4", "5", "6", "7", "8", "9"];
    this.currentPlayerId = 0;
  }
}

export default Game;
