-- +goose Up
BEGIN TRANSACTION;

    CREATE TABLE IF NOT EXISTS note (
        id BIGINT GENERATED ALWAYS AS IDENTITY,
        title      TEXT NOT NULL,
        text       TEXT NOT NULL,
        author     TEXT NOT NULL,
        email      TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT now(),
        updated_at TIMESTAMP NOT NULL,
        PRIMARY KEY(id)
    );

COMMIT;

-- +goose Down

BEGIN TRANSACTION;

    DROP TABLE note;

COMMIT;


