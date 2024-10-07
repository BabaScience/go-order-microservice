
api:
	go run cmd/api/main.go

worker:
	go run cmd/worker/main.go

test:
	go run cmd/test/main.go


up:
	docker-compose up -d

build-up:
	docker-compose up -d --build

down:
	docker-compose down

ps:
	docker-compose ps

prune:
	docker-compose down
	docker container prune -f
	docker image prune -f

postgres-up:
	docker-compose up -d postgres

api-up:
	docker-compose up -d api

worker-up:
	docker-compose up -d worker

frontend-up:
	docker-compose up -d frontend

log-api:
	docker-compose logs -f api

log-worker:
	docker-compose logs -f worker

log-postgres:
	docker-compose logs -f postgres

log-frontend:
	docker-compose logs -f frontend