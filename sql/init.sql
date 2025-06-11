
-- users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- books
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    user_id INTEGER REFERENCES users(id)
);

INSERT INTO users (name, password)
VALUES ('John Doe', 'mypw');


INSERT INTO books (name, content, author, user_id)
VALUES ('Book 1', 'This is the content of book 1', 'Author 1', 1);
