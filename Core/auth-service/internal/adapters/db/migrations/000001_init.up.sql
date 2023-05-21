CREATE TABLE users 
(
    login TEXT PRIMARY KEY,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    role VARCHAR(10) DEFAULT 'MEMBER'
);