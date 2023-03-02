class Player {
  id;
  username;
  password;

  constructor(copiedPlayer) {
    if (copiedPlayer) {
      this.id = copiedPlayer.id;
      this.username = copiedPlayer.username;
      this.password = copiedPlayer.password;
      return;
    }
    this.id = 0;
    this.username = "";
    this.password = "";
  }
}

export default Player;
