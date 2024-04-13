APP=banner-service
.PHONY: test

build:
	docker-compose build $(APP)

run:
	docker-compose up $(APP)

test:
	go test ./...
