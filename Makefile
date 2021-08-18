include .env

migrate:
	migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable -path ./migrations up

build:
	docker-compose build

run:
	docker-compose up -d

stop:
	docker-compose down

swag:
	swag init -g cmd/main.go