CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    owner_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    body TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT fk_notes_owner_id
       FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);