-- +goose Up
    CREATE TABLE IF NOT EXISTS note (
        id BIGINT GENERATED ALWAYS AS IDENTITY,
        title      TEXT NOT NULL,
        text       TEXT NOT NULL,
        author     TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT now(),
        updated_at TIMESTAMP,
        PRIMARY KEY(id)
    );

-- +goose Down
    DROP TABLE note;


