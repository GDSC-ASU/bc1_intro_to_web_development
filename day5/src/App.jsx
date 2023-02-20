import { useState } from "react";
import "./App.css";
import GameBoard from "./components/GameBoard";
import GameData from "./data/GameData";

function App() {
  const [gameData, setGameData] = useState(new GameData());

  return (
    <div>
      <GameBoard game={gameData} setGame={setGameData} />
    </div>
  );
}

export default App;
