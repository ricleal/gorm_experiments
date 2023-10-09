ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# start db
db-start:
	docker-compose -f $(ROOT_DIR)docker-compose.yaml up

db-stop:
	docker-compose -f $(ROOT_DIR)docker-compose.yaml down

db-cli:
	echo "Connecting to database..."
	mycli "$(DB_DSN_CLEAN)"

db-export-schema:
	mysqldump -u $(MYSQL_ROOT_USER) -h $(MYSQL_HOST) --port=$(MYSQL_PORT) --no-data \
		--password=$(MYSQL_ROOT_PASSWORD) $(MYSQL_DATABASE)
