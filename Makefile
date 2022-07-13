GO=go
SWAG=swag

.PHONY: docs
docs:
	$(SWAG) init

.PHONY: run
run:
	docker-compose up -d api_client
	docker-compose up -d api_docs

.PHONE: stop
stop:
	docker-compose down
