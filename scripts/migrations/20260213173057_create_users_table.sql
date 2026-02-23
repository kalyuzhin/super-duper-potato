-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         BIGINT  NOT NULL,
    auth_key   BLOB    NOT NULL,
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
    created_at INTEGER NOT NULL DEFAULT (strftime('%s', 'now'))
);

CREATE UNIQUE INDEX users_id_idx ON users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
