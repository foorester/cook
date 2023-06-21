CREATE TABLE cook.users (
                          id UUID PRIMARY KEY,
                          username VARCHAR(255) NOT NULL,
                          email VARCHAR(255) NOT NULL,
                          password TEXT NOT NULL,
                          name VARCHAR(255) NOT NULL
);
