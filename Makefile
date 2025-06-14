.PHONY: build
build:
	go build -o ./build/siloamteens_app ./cmd/main.go

.PHONY: run
run:
	./build/siloamteens_app

.PHONY: swag
swag:
	swag init -g cmd/main.go
