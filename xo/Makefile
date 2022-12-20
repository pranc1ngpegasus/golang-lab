.PHONY: migrate.plan
migrate.plan:
	@go run github.com/k0kubun/sqldef/cmd/psqldef -U${DATABASE_USER} -W${DATABASE_PASSWORD} -h${DATABASE_HOST} -p${DATABASE_PORT} ${DATABASE_NAME} --dry-run < schema.sql

.PHONY: migrate.apply
migrate.apply:
	@go run github.com/k0kubun/sqldef/cmd/psqldef -U${DATABASE_USER} -W${DATABASE_PASSWORD} -h${DATABASE_HOST} -p${DATABASE_PORT} ${DATABASE_NAME} < schema.sql

.PHONY: middleware.build
middleware.build:
	docker compose -f docker/docker-compose.yml build

.PHONY: middleware.up
middleware.up:
	docker compose -f docker/docker-compose.yml up -d

.PHONY: middleware.down
middleware.down:
	docker compose -f docker/docker-compose.yml stop

.PHONY: middleware.logs
middleware.logs:
	docker compose -f docker/docker-compose.yml logs -f
