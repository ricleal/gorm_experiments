# include common Makefile
include ../Makefile

### DB migration targets

# https://github.com/golang-migrate/migrate
# brew install golang-migrate
db-migrate-up: ## Run database upgrade migrations
	migrate -verbose -database "$(DB_DSN)" -path migrations up

db-migrate-down:  ## Run database downgrade the last migration
	migrate -verbose -database "$(DB_DSN)" -path migrations down 1

db-migrate-version:  ## Print the current migration version
	migrate -verbose -database "$(DB_DSN)" -path migrations version

db-migrate-create:  ## Create a new migration file
	@if [ -z "$(name)" ]; then echo "name is required"; exit 1; fi
	migrate create -ext sql -dir migrations -seq $(name)

db-migrate-force:  ## Force mark the migration version
	@if [ -z "$(version)" ]; then echo "version is required"; exit 1; fi
	migrate -verbose -database "$(DB_DSN)" -path migrations force $(version)

### Run targets
run: ## Run the application
	@go run main.go
