run:
	go run main.go

build:
	docker build -t goapi_development .

up:
	docker compose up -d

down:
	docker compose down
