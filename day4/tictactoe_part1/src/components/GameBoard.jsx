import GameData from "../data/GameData";
import "./GameBoard.css"
import Stats from "./Stats";

function GameBoard({ game, setGame }) {

    const play = (cellNumber) => {
        game.markCell(cellNumber);
        game.checkWinner();
        setGame(new GameData(game));
    }

  return (
    <div>
        <div>
            <Stats xScore={game.xScore} oScore={game.oScore} turn={game.turn}/>
        </div>
      {game.cells.map((cell, index) => {
        return (
          <div className="cell" key={index}>
            <button className="cell" onClick={() => {play(index)}}>{cell}</button>
            {((index + 1) % 3 === 0) && <br />}
          </div>
        );
      })}
    </div>
  );
}

export default GameBoard;
