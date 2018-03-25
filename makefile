.PHONY: all db_clean

db_clean:
	@docker rm shjp_data && echo "Data cleanup complete." || echo "No data container to cleanup"
	@docker stop shjp_db && docker rm shjp_db && echo "DB cleanup complete." || echo "No DB container. Nothing to do."

db_init: db_clean
	@docker create -v /shjp_data --name shjp_data postgres
	@docker run --name shjp_db --volumes-from shjp_data -e POSTGRES_USER=shjp -e POSTGRES_PASSWORD=hellochurch -e POSTGRES_DB=shjp_youth -p 5432:5432 -d postgres
	@echo "DB container initialized"

db_migrate_dev:
	@goose -dir db/migrations postgres "user=shjp password=hellochurch host=localhost port=5432 dbname=shjp_youth sslmode=disable" up

db: db_clean db_init db_migrate_dev

db_fixtures:
	@go run fixtures/main.go

db_reset:
	@goose -dir db/migrations postgres "user=shjp password=hellochurch host=localhost port=5432 dbname=shjp_youth sslmode=disable" down
	@goose -dir db/migrations postgres "user=shjp password=hellochurch host=localhost port=5432 dbname=shjp_youth sslmode=disable" up
	@make db_fixtures

run:
	@go run main.go