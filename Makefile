APP=banner-service
.PHONY: test

build:
	docker-compose build $(APP)

run:
	docker-compose up -d $(APP)

stop:
	docker-compose down

test:
	make run
	go test ./test/ -count=1
	make stop

linter:
	golangci-lint run ./... --config=./golangci.yaml
