.DEFAULT_GOAL := build

env:
	@echo "Set up enviroment"
	@go get
	@go mod tidy

build: env
	@echo "Make build"
	@mkdir -p cache/content cache/meta
	@go run ./main.go
	@cd ./cache; zip -r ../cache.zip .; cd ..;