# Nome da imagem
IMAGE_NAME = goapi_development

run:
	go run main.go

# Build local
build:
	go build -o main .

# Build da imagem Docker
docker-build:
	docker build -t $(IMAGE_NAME) .

up:
	docker compose up -d

down:
	docker compose down

# Reinicia tudo
restart: 
	down up

# Mostra logs da API em tempo real
logs:
	docker compose logs -f api

rebuild: down
	docker compose up -d --build

# Deploy para servidor (Exemplo usando SSH e Docker)
deploy1:
	docker save $(IMAGE_NAME) | ssh user@server "docker load && docker compose up -d --build"

deploy2:
	docker build -t usuario/goapi:latest .
	docker push usuario/goapi:latest