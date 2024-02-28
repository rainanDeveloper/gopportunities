.PHONY: default run run-with-docs build test docs clean

APP_NAME=gopportunities
DIST_FOLDER=dist

default: run-with-docs

run: 
	@go run main.go

run-with-docs:
	@swag init
	@go run main.go

build:
	@go build -o ${DIST_FOLDER}/${APP_NAME} main.go

test:
	@go test ./ ...

docs:
	@swag init

clean:
	@rm -rf ${DIST_FOLDER}
	@rm -rf docs