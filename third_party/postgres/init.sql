CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    done BOOL DEFAULT FALSE
);