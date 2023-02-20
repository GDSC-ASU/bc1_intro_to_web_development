import GameData from "../data/GameData";
import "./GameBoard.css";
import Stats from "./Stats";

function GameBoard({ game, setGame }) {
  const play = (cellNumber) => {
    game.markCell(cellNumber);
    game.checkWinner();
    setGame(new GameData(game));
  };

  return (
    <div className="bigContainer">
      <div className="board">
        {game.cells.map((cell, index) => {
          return (
            <div className="cell" key={index}>
              <button
                className="cell"
                onClick={() => {
                  play(index);
                }}
              >
                {(cell === "X" || cell === "O") && cell}
                {cell !== "X" && cell !== "O" && " "}
              </button>
              {(index + 1) % 3 === 0 && <br />}
            </div>
          );
        })}
      </div>
      <div>
        <Stats xScore={game.xScore} oScore={game.oScore} turn={game.turn} />
      </div>
    </div>
  );
}

export default GameBoard;
