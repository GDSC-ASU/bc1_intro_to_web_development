CREATE DATABASE ttt;

CREATE TABLE IF NOT EXISTS ttt.players (
    id INTEGER AUTO_INCREMENT,
    username VARCHAR(60) UNIQUE NOT NULL,
    password VARCHAR(200),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS ttt.sessions (
    id INTEGER AUTO_INCREMENT,
    username VARCHAR(60),
    token VARCHAR(32) UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS ttt.game (
    id INTEGER AUTO_INCREMENT,
    x_player_id INTEGER,
    o_player_id INTEGER,
    x_score INTEGER,
    o_score INTEGER,
    cells VARCHAR(20),
    current_player_id INTEGER,
    PRIMARY KEY (id),
    FOREIGN KEY (x_player_id) REFERENCES ttt.players(id),
    FOREIGN KEY (current_player_id) REFERENCES ttt.players(id)
);
