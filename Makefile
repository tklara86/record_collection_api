include .env

server:
	go run ./cmd/api

webpack:
	npm run watch

# migrate create -ext sql -dir ./migrations -seq init_schema
migrateup:
	migrate -path migrations -database "$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)" up

migratedown:
	migrate -path migrations -database "$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)" down
