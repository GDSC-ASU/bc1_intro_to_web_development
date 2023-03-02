import Game from "../models/Game";

class GameRequests {
  static async newGame() {
    return await fetch("http://localhost:8080/game/new", {
      method: "GET",
      mode: "cors",
      headers: {
        Authorization: localStorage.getItem("token"),
      },
    })
      .then((resp) => {
        return resp.json();
      })
      .then((game) => {
        return new Game(game);
      })
      .catch((err) => console.log(err));
  }

  static async joinGame(gameId) {
    return await fetch(`http://localhost:8080/game/join?id=${gameId}`, {
      method: "GET",
      mode: "cors",
      headers: {
        Authorization: localStorage.getItem("token"),
      },
    })
      .then((resp) => {
        return resp.json();
      })
      .then((game) => {
        return new Game(game);
      })
      .catch((err) => console.log(err));
  }

  static async playTurn(reqBody) {
    return fetch("http://localhost:8080/game/play", {
      method: "POST",
      mode: "cors",
      headers: {
        Authorization: localStorage.getItem("token"),
      },
      body: JSON.stringify(reqBody),
    });
  }

  static async getGame(id) {
    return await fetch(`http://localhost:8080/game/get?id=${id}`, {
      method: "GET",
      mode: "cors",
      headers: {
        Authorization: localStorage.getItem("token"),
      },
    })
      .then((resp) => {
        return resp.json();
      })
      .then((game) => {
        return new Game(game);
      })
      .catch((err) => console.log(err));
  }
}

export default GameRequests;
