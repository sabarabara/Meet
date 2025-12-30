CREATE TABLE IF NOT EXISTS RECRUIT_READ (
    recruitid UUID PRIMARY KEY REFERENCES RECRUIT(recruitid) ON DELETE CASCADE,
    date DATE,
    area VARCHAR(255),
    vacant_man INT,
    vacant_woman INT
);

CREATE INDEX IF NOT EXISTS idx_recruit_read_date ON RECRUIT_READ(date);
CREATE INDEX IF NOT EXISTS idx_recruit_read_area ON RECRUIT_READ(area);
CREATE INDEX IF NOT EXISTS idx_recruit_read_vacant_man ON RECRUIT_READ(vacant_man);
CREATE INDEX IF NOT EXISTS idx_recruit_read_vacant_woman ON RECRUIT_READ(vacant_woman);
