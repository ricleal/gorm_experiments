
# DB_DSN = ${DB_DSN}

# start postgresql
db-start:
	@docker-compose -f docker-compose.yaml up

db-stop:
	@docker-compose -f docker-compose.yaml down

db-cli:
	echo "Connecting to database..."
	@mycli $(DB_DSN)

### DB migration targets

# https://github.com/golang-migrate/migrate
# brew install golang-migrate
# db-migrate-up: ## Run database upgrade migrations
# 	migrate -verbose -database $(DB_DSN) -path migrations up

# db-migrate-down:  ## Run database downgrade the last migration
# 	migrate -verbose -database $(DB_DSN) -path migrations down 1

# db-migrate-version:  ## Print the current migration version
# 	migrate -verbose -database $(DB_DSN) -path migrations version
