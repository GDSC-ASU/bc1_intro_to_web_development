import GameData from "../data/GameData";
import Game from "../models/Game";
import GameRequests from "../utils/GameRequests";
import LoginRequests from "../utils/LoginRequests";
import "./GameBoard.css";
import Stats from "./Stats";
import { useEffect } from "react";
import { useState } from "react";
import Player from "../models/Player";

function GameBoard({ game, setGame }) {
  const [player, setPlayer] = useState(new Player());

  const play = async (cellNumber) => {
    await GameRequests.playTurn({
      cellNumber: cellNumber,
      gameId: game.id,
    }).then(async (resp) => {
      if (resp.ok) {
        setGame(new Game(await resp.json()));
        return;
      }
      window.alert(await resp.text());
    });
  };

  useEffect(() => {
    (async () => {
      setPlayer(await LoginRequests.tokenLogin());
    })();

    const interval = setInterval(async () => {
      setGame(await GameRequests.getGame(game.id));
    }, 1000);

    return () => clearInterval(interval);
  }, []);

  const getCurrentPlayer = () => {
    switch (game.currentPlayerId) {
      case game.xPlayerId:
        return "X";
      case game.oPlayerId:
        return "O";
    }
    return "";
  };

  const isPlayerPlaying = () => {
    return player.id === game.currentPlayerId;
  };

  return (
    <div className="bigContainer">
      <h2 className="playerStatus">
        {isPlayerPlaying() ? "Your turn" : "Wait for your opponent"}
      </h2>
      <div className="board">
        {game.cells.map((cell, index) => {
          return (
            <div className="cell" key={index}>
              <button
                className="cell"
                onClick={async () => {
                  await play(index);
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
        <Stats
          xScore={game.xScore}
          oScore={game.oScore}
          turn={getCurrentPlayer()}
          gameId={game.id}
        />
      </div>
    </div>
  );
}

export default GameBoard;
