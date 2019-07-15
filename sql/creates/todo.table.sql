CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    listId INT REFRENCES list(id),
    name VARCHAR(150),
    isComplete BIT,
    created TIMESTAMP,
    deleted TIMESTAMP,
    order INT
)