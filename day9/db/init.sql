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
    token VARCHAR(32),
    PRIMARY KEY (id)
);