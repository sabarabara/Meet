CREATE TABLE IF NOT EXISTS RECRUIT (
    recruitid UUID PRIMARY KEY,
    userid UUID NOT NULL REFERENCES USERS(userid) ON DELETE CASCADE,
    area VARCHAR(255),
    imgurl VARCHAR(255),
    man INT,
    woman INT,
    vacant_man INT,
    vacant_woman INT,
    comment TEXT,
    date DATE
);
