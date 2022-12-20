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
