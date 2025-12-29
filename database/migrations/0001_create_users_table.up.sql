CREATE TABLE IF NOT EXISTS USERS (
    userid UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    imgurl VARCHAR(255),
    pronunciation VARCHAR(255),
    selfintroduce TEXT,
    stars INT
);
