import Player from "../models/Player";

class LoginRequests {
  static async login(player) {
    return await fetch("http://localhost:8080/player/login", {
      method: "POST",
      mode: "cors",
      body: JSON.stringify(player),
    })
      .then((resp) => {
        return resp.json();
      })
      .then((data) => {
        localStorage.setItem("token", data["token"]);
      })
      .catch((err) => console.log(err));
  }

  static async signup(player) {
    return await fetch("http://localhost:8080/player/signup", {
      method: "POST",
      mode: "cors",
      body: JSON.stringify(player),
    })
      .then((resp) => {})
      .catch((err) => console.log(err));
  }

  static async tokenLogin() {
    return await fetch("http://localhost:8080/player/token", {
      method: "GET",
      mode: "cors",
      headers: { Authorization: localStorage.getItem("token") },
    })
      .then((resp) => {
        return resp.json();
      })
      .then((player) => {
        return new Player(player);
      })
      .catch((err) => console.log(err));
  }
}

export default LoginRequests;
