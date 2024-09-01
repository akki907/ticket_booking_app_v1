start: build
	@./bin/main

build:
	@go build -o ./bin ./cmd/api/main.go

docker_start: 
	@docker compose up --build
	
docker_stop:
	@docker-compose rm -v --force --stop
	@docker rmi ticket-booking
