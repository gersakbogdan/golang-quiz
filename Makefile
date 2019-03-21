.PHONY: build proto run stop

build:
	@docker-compose build
	@docker-compose up -d --remove-orphans

run:
	@docker-compose exec quizer quizer client

stop:
	@docker-compose stop

proto:
	@for f in proto/*.proto; do \
		protoc --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done
