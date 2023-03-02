import { useEffect, useState } from "react";
import GameRequests from "../utils/GameRequests";
import { default as GameData } from "../models/Game";
import GameBoard from "../components/GameBoard";
import LoginRequests from "../utils/LoginRequests";
import { useNavigate } from "react-router-dom";

function Game() {
  const [gameData, setGameData] = useState(new GameData());

  const [joinGameId, setJoinGameId] = useState(0);

  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      const player = await LoginRequests.tokenLogin();
      if (!player || !player.id) {
        navigate("/login");
      }
    })();
  }, []);

  const joinGameIdOnChangeHandler = (event) => {
    setJoinGameId(parseInt(event.target.value));
  };

  const joinGame = async () => {
    setGameData(await GameRequests.joinGame(joinGameId));
  };

  const newGame = async () => {
    setGameData(await GameRequests.newGame());
  };

  return (
    <div>
      {gameData.id === 0 && (
        <div>
          <input
            placeholder="Game Id"
            type="number"
            value={joinGameId}
            onChange={joinGameIdOnChangeHandler}
          />
          <button onClick={joinGame}>Join Game</button>
          <br />
          <button onClick={newGame}>New Game</button>
        </div>
      )}
      {gameData.id !== 0 && (
        <div>
          <GameBoard game={gameData} setGame={setGameData} />
        </div>
      )}
    </div>
  );
}

export default Game;
