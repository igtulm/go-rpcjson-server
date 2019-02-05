include app.env
export

SHELL := /bin/bash
TARGET_BINARY = bin/app

.PHONY: build clean run

all: build

deps:
	go get github.com/kelseyhightower/envconfig
	go get github.com/lib/pq
	go get github.com/gorilla/mux
	go get github.com/gorilla/rpc
	go get github.com/gorilla/rpc/json

db:
	echo "DROP DATABASE IF EXISTS $(APP_DB_DATABASE); CREATE DATABASE $(APP_DB_DATABASE);" | PGPASSWORD=$(APP_DB_PASSWORD) psql postgres -U $(APP_DB_USER) -h $(APP_DB_HOST) -w
	echo "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";" | PGPASSWORD=$(APP_DB_PASSWORD) psql $(APP_DB_NAME) -U $(APP_DB_USER) -h $(APP_DB_HOST) -w
	echo "CREATE TABLE users(\
            id UUID DEFAULT uuid_generate_v4(),\
            login TEXT,\
            created_at TIMESTAMP,\
            PRIMARY KEY(id)\
    );"| PGPASSWORD=$(APP_DB_PASSWORD) psql $(APP_DB_NAME) -U $(APP_DB_USER) -h $(APP_DB_HOST) -w
	echo "CREATE UNIQUE INDEX login_index ON users (login);"

build:
	@go build -o $(TARGET_BINARY)

clean:
	@rm -rf $(TARGET_BINARY)

run:
	@$(TARGET_BINARY)

test:
	@go test
