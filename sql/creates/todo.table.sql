CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    listId INT REFERENCES todoList(id),
    name VARCHAR(150),
    isComplete BIT,
    created TIMESTAMP,
    deleted TIMESTAMP,
    sort INT
)