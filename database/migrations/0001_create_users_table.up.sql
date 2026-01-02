CREATE TABLE IF NOT EXISTS USERS (
    userid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) NOT NULL,
    imgurl VARCHAR(255),
    pronunciation VARCHAR(255),
    selfintroduce TEXT,
    stars FLOAT DEFAULT 0.0
);
