-- +goose Up
-- +goose StatementBegin
CREATE TABLE vault
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id        BIGINT  NOT NULL,
    service        TEXT    NOT NULL,
    login          BLOB    NOT NULL,
    login_nonce    BLOB    NOT NULL,
    password       BLOB    NOT NULL,
    password_nonce BLOB    NOT NULL,
    secret         BLOB,
    secret_nonce   BLOB,
    created_at     INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
    updated_at     INTEGER NOT NULL DEFAULT (strftime('%s', 'now'))
);

CREATE UNIQUE INDEX vault_user_id_service_idx ON vault(user_id, service);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vault;
-- +goose StatementEnd
