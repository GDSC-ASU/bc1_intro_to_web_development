import { useState } from "react";
import Player from "../models/Player";
import LoginRequests from "../utils/LoginRequests";
import { useNavigate } from "react-router-dom";

function Login() {
  const navigate = useNavigate();
  const [player, setPlayer] = useState(new Player());

  const usernameOnChangeHandler = (event) => {
    player.username = event.target.value;
    setPlayer(new Player(player));
  };

  const passwordOnChangeHandler = (event) => {
    player.password = event.target.value;
    setPlayer(new Player(player));
  };

  const login = async () => {
    await LoginRequests.login(player);
    console.log(player);
    navigate("/");
  };

  return (
    <div>
      <input
        type="text"
        placeholder="Username"
        value={player.username}
        onChange={usernameOnChangeHandler}
      />
      <br />
      <input
        type="password"
        placeholder="Password"
        value={player.password}
        onChange={passwordOnChangeHandler}
      />
      <br />
      <button onClick={login}>Login</button>
    </div>
  );
}

export default Login;
