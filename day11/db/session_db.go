package db

import "tictactoe/models"

type SessionDB struct{}

func (s *SessionDB) Create(session models.Session) error {
	_, err := GetInstance().Exec("INSERT INTO ttt.sessions (username, token) VALUES (?, ?)", session.Username, session.Token)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionDB) Get(token string) (models.Session, error) {
	var session models.Session

	err := GetInstance().
		QueryRow("SELECT username, token FROM ttt.sessions WHERE token=?", token).
		Scan(&session.Username, &session.Token)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}
