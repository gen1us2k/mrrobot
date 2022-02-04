.PHONY: help

help:
	        @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## Runs linter against the code
	        golangci-lint run ./...

test: ## Run tests locally
	        go test ./...

build_docker: ## Build docker image
	        docker build -t mrrobot .

build_linux: ## Build executable for linux system
	        GOOS=linux GOARCH=amd64 go build -o mrrobot cmd/mrrobot/main.go

zip: build_linux  ## Build and create a zip archive for deploying to AWS lambda
	        zip main.zip mrrobot
