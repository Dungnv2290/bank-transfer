ENVIRONMENT=development
SYSTEM=bank-transfer
SYSTEM_VERSION=$(shell git branch --show-current | cut -d '/' -f2)
PWD=$(shell pwd -L)
DOCKER_RUN=docker run --rm -it -w /app -v ${PWD}:/app -v ${GOPATH}/pkg/mod/cache:/go/pkg/mod/cache golang:1.16-buster

.PHONY: all
all: help
help: ## Display help screen
	@echo "Usage:"
	@echo " make [COMMAND]"
	@echo " make help \n"
	@echo "Commands: \n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Create environment variables
	cp .env.example .env

.PHONY: test
test: ## Run golang tests
	${DOCKER_RUN} go test -cover -race ./...

.PHONY: test-local
test-local: ## Run local golang tests
	go test -cover -race ./...

.PHONY: up
up: ## Run docker-compose up for creating and starting containers
	docker-compose up -d

.PHONY: test-report
test-report: ## Run tests with HTML coverage report
	${DOCKER_RUN} go test -covermode=count -coverprofile coverage.out -p=1 ./... && \
	go tool cover -html=coverage.out -o coverage.html && \
	xdg-open ./coverage.html

.PHONY: test-report-function
test-report-function: ## Run tests with function report -covermode=set
	${DOCKER_RUN} go test -covermode=set -coverprofile=coverage.out -p=1 ./... && \
	go tool cover -func=coverage.out

.PHONY: logs
logs: ## View container log
	docker-compose logs -f app