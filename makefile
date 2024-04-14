# LOG_DIR=./logs
# SWAG_DIRS=./internal/app/delivery/http/v1/,./internal/banner/delivery/http/v1/handlers,./internal/banner/delivery/http/v1/models/request,./internal/banner/delivery/http/v1/models/response,./external/auth/delivery/http/v1/handlers,./internal/app/delivery/http/tools
# include ./config/env/api_test.env
# export $(shell sed 's/=.*//' ./config/env/api_test.env)

# Запуск

.PHONY: run
run:
	sudo docker compose --env-file ./docker_run.env up -d

.PHONY: run-verbose
run-verbose:
	sudo docker compose --env-file ./docker_run.env up

.PHONY: stop
stop:
	sudo docker compose --env-file ./docker_run.env stop

.PHONY: down
down:
	sudo docker compose --env-file ./docker_run.env down

# Сборка

.PHONY: build-banner
build-banner:
	go build -o server -v ./cmd/banner

.PHONY: build-docker-banner
build-docker-banner:
	sudo docker build --no-cache --network host -f ./docker/banner.Dockerfile . --tag banner

.PHONY: build-docker-all
build-docker-all: build-docker-banner

.PHONY: fmt
fmt:
	gofumpt -e -w -extra .
	goimports -e -w .