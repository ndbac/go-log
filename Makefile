DB_URI=postgresql://golog:golog@localhost:5432/golog?sslmode=disable

postgres:
	docker run --name postgres --network golog -p 5432:5432 -e POSTGRES_USER=golog -e POSTGRES_PASSWORD=golog -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=golog --owner=golog golog

dropdb:
	docker exec -it postgres dropdb golog

migrate-up:
	migrate -path db/migrations -database "$(DB_URI)" -verbose up

migrate-up-1:
	migrate -path db/migrations -database "$(DB_URI)" -verbose up 1

migrate-down:
	migrate -path db/migrations -database "$(DB_URI)" -verbose down

migrate-down-1:
	migrate -path db/migrations -database "$(DB_URI)" -verbose down 1

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

.PHONY: postgres createdb dropdb migrate-up migrate-up-1 migrate-down migrate-down-1