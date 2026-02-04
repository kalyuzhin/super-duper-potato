include .env
export $(shell sed 's/=.*//' .env)

DB_SETUP := user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable
SQLITE_FILE = ./passwords.db
MIGRATIONS_DIR = scripts/migrations

.PHONY: help
## prints help about all targets
help:
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@awk '                                \
		BEGIN { comment=""; }             \
		/^\s*##/ {                         \
		    comment = substr($$0, index($$0,$$2)); next; \
		}                                  \
		/^[a-zA-Z0-9_-]+:/ {               \
		    target = $$1;                  \
		    sub(":", "", target);          \
		    if (comment != "") {           \
		        printf "  %-17s %s\n", target, comment; \
		        comment="";                \
		    }                              \
		}' $(MAKEFILE_LIST)
	@echo ""

.PHONY: build
build:
	go build -o bin/main cmd/passwordmanager/main.go

.PHONY: migrate-pg
migrate-pg:
	goose -dir "$(MIGRATIONS_DIR)" postgres "$(DB_SETUP)" up

.PHONY: migrate-sqlite
migrate-sqlite:
	goose -dir "$(MIGRATIONS_DIR)" sqlite "$(SQLITE_FILE)" up

.PHONY: migrate-down-sqlite
migrate-down-sqlite:
	goose -dir "$(MIGRATIONS_DIR)" sqlite "$(SQLITE_FILE)" down