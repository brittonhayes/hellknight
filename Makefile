.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build:
	docker build . -t brittonhayes/hellknight:latest

.PHONY: deploy
deploy:
	railway up