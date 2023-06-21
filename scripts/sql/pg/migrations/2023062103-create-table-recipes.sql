CREATE TABLE cook.recipes (
                        id UUID PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        description VARCHAR(255) NOT NULL,
                        book_id UUID NOT NULL
);
