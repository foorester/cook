CREATE TABLE cook.books (
                        id UUID PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        description VARCHAR(255) NOT NULL,
                        owner_id UUID NOT NULL
);
