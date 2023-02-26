package db

import "tictactoe/models"

type PlayerDB struct{}

func (p *PlayerDB) Create(player models.Player) (models.Player, error) {
	result, err := GetInstance().Exec(
		"INSERT INTO ttt.players (username, password, created_at) VALUES (?, ?, current_timestamp)",
		player.Username, player.Password)
	if err != nil {
		return models.Player{}, err
	}

	playerId, err := result.LastInsertId()
	if err != nil {
		return models.Player{}, err
	}

	player.Id = int(playerId)

	return player, nil
}

func (p *PlayerDB) Get(username string) (models.Player, error) {
	var player models.Player
	err := GetInstance().
		QueryRow("SELECT id, username, password FROM ttt.players WHERE username=?", username).
		Scan(&player.Id, &player.Username, &player.Password)

	if err != nil {
		return models.Player{}, err
	}

	return player, nil
}

func (p *PlayerDB) Update(username string, player models.Player) error {
	_, err := GetInstance().Exec("UPDATE ttt.players SET username=?, password=? WHERE username=?", player.Username, player.Password, username)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlayerDB) Delete(username string) error {
	_, err := GetInstance().Exec("DELETE FROM ttt.players WHERE username=?", username)
	if err != nil {
		return err
	}
	return nil
}