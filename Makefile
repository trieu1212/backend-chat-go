include .env
up:	
	@echo "Starting up services"
	docker-compose up -d --remove-orphans

down:
	@echo "Stopping services"
	docker-compose down

build:
	go build -o ${TRIEU} ./cmd/api/main.go
	
start:
	./${TRIEU}

restart: build start