--UP
CREATE TABLE recipes (
                         id TEXT PRIMARY KEY,
                         name TEXT NOT NULL,
                         description TEXT NOT NULL,
                         book_id TEXT NOT NULL,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--DOWN
DROP TABLE recipes;
