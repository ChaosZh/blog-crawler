.DEFAULT_GOAL := build

vet:
	go vet ./...

fmt:
	go fmt ./...

env:
	@echo "Set up enviroment"
	@go get
	@go mod tidy

build: vet fmt env
	@echo "Make build"
	@mkdir -p cache/content cache/meta
	@go run ./main.go
	@cd ./cache; zip -r ../cache.zip .; cd ..;