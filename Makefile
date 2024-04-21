.PHONY: help

include .env
export $(shell sed 's/=.*//' .env)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

install-swag: ## Install go-swagger to generate docs
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: install-swag

run: ## Run server locally
	go run cmd/main.go

.PHONY: run

docs: ## Generate docs
	swag init -g ./cmd/main.go -o cmd/docs

.PHONY: docs

test-api: ## Test pkg/api
	go test translator/pkg/api

.PHONY: test-api