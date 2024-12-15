CREATE TABLE users (
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
);

CREATE TABLE quizzes (
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    title TEXT NOT NULL,
    questions TEXT NOT NULL,
    results TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    time_to_live TIMESTAMP DEFAULT (datetime ('now', '+1 day')),
    link_to_quiz TEXT NOT NULL,
    user_id TEXT NOT NULL
);