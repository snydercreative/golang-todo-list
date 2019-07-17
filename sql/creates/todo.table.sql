DROP TABLE IF EXISTS todo;

CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    listId INT REFERENCES todoList(id),
    name VARCHAR(150),
    isComplete BOOLEAN DEFAULT FALSE,
    created TIMESTAMP,
    deleted TIMESTAMP NULL,
    sort INT
);