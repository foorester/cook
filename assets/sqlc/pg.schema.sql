-- Create the 'cook' schema
CREATE SCHEMA cook;

-- Create the 'users' table
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

-- Create the 'books' table
CREATE TABLE books (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    owner_id UUID NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users (id)
);

-- Create the 'recipes' table
CREATE TABLE recipes (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    book_id UUID NOT NULL,
    FOREIGN KEY (book_id) REFERENCES books (id)
);

-- Create the 'ingredients' table
CREATE TABLE ingredients (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    recipe_id UUID NOT NULL,
    qty TEXT NOT NULL,
    unit TEXT NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes (id)
);

-- Create the 'steps' table
CREATE TABLE steps (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    recipe_id UUID NOT NULL,
    duration TEXT NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes (id)
);

