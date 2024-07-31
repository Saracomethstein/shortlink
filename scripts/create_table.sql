CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL,
    short_url VARCHAR(50) UNIQUE NOT NULL,
    original_url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users_log (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL,
    session_id TEXT NOT NULL
);
