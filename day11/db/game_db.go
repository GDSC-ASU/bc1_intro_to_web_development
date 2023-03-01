package db

import "tictactoe/models"

type GameDB struct{}

func (g *GameDB) Create(game models.Game) (models.Game, error) {
	result, err := GetInstance().Exec(
		"INSERT INTO ttt.games (x_player_id, o_player_id, x_score, o_score, current_player_id, cells) VALUES (?, ?, ?, ?, ?, ?)",
		game.XPlayerId, 0, 0, 0, game.XPlayerId, "1 2 3 4 5 6 7 8 9")
	if err != nil {
		return models.Game{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return models.Game{}, err
	}

	game.Id = int(lastId)
	game.Cells = "1 2 3 4 5 6 7 8 9"
	game.CurrentPlayerId = game.XPlayerId
	game.XScore = 0
	game.OScore = 0

	return game, nil
}

func (g *GameDB) Get(id int) (models.Game, error) {
	var game models.Game
	err := GetInstance().
		QueryRow("SELECT * FROM ttt.games WHERE id=?", id).
		Scan(&game.Id, &game.XPlayerId, &game.OPlayerId, &game.XScore, &game.OScore,
			&game.Cells, &game.CurrentPlayerId)
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}

func (p *GameDB) Update(id int, game models.Game) error {
	_, err := GetInstance().Exec(
		"UPDATE ttt.games SET cells=?, x_score=?, o_score=?, current_player_id=?, o_player_id=?",
		game.Cells, game.XScore, game.OScore, game.CurrentPlayerId, game.OPlayerId)
	if err != nil {
		return err
	}
	return nil
}

func (p *GameDB) Delete(id int) error {
	_, err := GetInstance().Exec("DELETE FROM ttt.games WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
