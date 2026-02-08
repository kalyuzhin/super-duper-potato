-- +goose Up
-- +goose StatementBegin
CREATE TABLE meta
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT    NOT NULL UNIQUE,
    kdf_type   TEXT    NOT NULL,
    salt       BLOB    NOT NULL,
    time       INT     NOT NULL,
    threads    INT     NOT NULL,
    memory     INT     NOT NULL,
    key_length INT     NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s', 'now'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE meta;
-- +goose StatementEnd
