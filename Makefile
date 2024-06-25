.PHONY: build
build:
	go build -o cmd/main.go

.PHONY: run
run:
	go run cmd/main.go

.PHONY: swagger
swagger:
	swag init -g cmd/main.go \
	--exclude ./internal/ \
	-o ./docs/
