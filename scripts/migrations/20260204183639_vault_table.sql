-- +goose Up
-- +goose StatementBegin
CREATE TABLE vault
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    service    TEXT    NOT NULL,
    login      BLOB    NOT NULL,
    nonce      BLOB    NOT NULL,
    password   BLOB    NOT NULL,
    secret     BLOB,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s', 'now'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vault;
-- +goose StatementEnd
