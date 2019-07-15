CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150),
    created TIMESTAMP,
    deleted TIMESTAMP
)