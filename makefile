.DEFAULT_GOAL := build

env:
	@echo "Set up enviroment"
	@go get
	@go mod tidy

build: env
	@echo "Make build"
	@go run ./main.go