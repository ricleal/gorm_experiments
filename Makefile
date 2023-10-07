
# start postgresql
db-start:
	@docker-compose -f docker-compose.yaml up

db-stop:
	@docker-compose -f docker-compose.yaml down

db-cli:
	echo "Connecting to database..."
	@mycli $(DB_DSN)

db-export-schema:
	mysqldump -u $(MYSQL_ROOT_USER) -h $(MYSQL_HOST) --port=$(MYSQL_PORT) --no-data \
		--password=$(MYSQL_ROOT_PASSWORD) $(MYSQL_DATABASE)

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
