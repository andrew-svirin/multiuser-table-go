DOCKER_DIR := infrastructure
DC := docker-compose

SERVER_CONTAINER := server-multiuser-table-go
SERVER_CONTAINER_EXEC := $(DC) exec $(SERVER_CONTAINER)

DB_CONTAINER := db-multiuser-table-go
DB_CONTAINER_EXEC := $(DC) exec $(DB_CONTAINER)

BASH := /bin/bash
GO := /usr/local/go/bin/go

build:
	cd $(DOCKER_DIR) && $(DC) build

build-no-cache:
	cd $(DOCKER_DIR) && $(DC) build --no-cache

start:
	cd $(DOCKER_DIR) && $(DC) up -d

stop:
	cd $(DOCKER_DIR) && $(DC) down

restart: stop start

bash-server:
	cd $(DOCKER_DIR) && $(SERVER_CONTAINER_EXEC) $(BASH)

bash-db:
	cd $(DOCKER_DIR) && $(DB_CONTAINER_EXEC) $(BASH)

serve-server:
	cd $(DOCKER_DIR) && $(SERVER_CONTAINER_EXEC) $(GO) run .
