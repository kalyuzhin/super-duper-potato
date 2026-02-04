-- +goose Up
-- +goose StatementBegin
CREATE TABLE meta
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    name          TEXT    NOT NULL,
    salt          BLOB    NOT NULL,
    argon_time    INT     NOT NULL,
    argon_threads INT     NOT NULL,
    argon_memory  INT     NOT NULL,
    key_length    INT     NOT NULL,
    created_at    INTEGER NOT NULL DEFAULT (strftime('%s', 'now'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE meta;
-- +goose StatementEnd
