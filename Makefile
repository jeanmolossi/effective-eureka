GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test
SWAG=swag
CODE_PATH=./src

.PHONY: docs
docs:
	$(SWAG) init

.PHONY: run
run:
	$(SWAG) init
	docker-compose up -d api_db
	docker-compose up -d api_client
	docker-compose up -d api_docs

.PHONE: stop
stop:
	docker-compose down

.PHONY: mock
mock:
	./local-scripts/install-mockery
	source ~/.bashrc
	mockery --all

.PHONY: test
test:
	cd $(CODE_PATH) && mkdir -p ./tests
	cd $(CODE_PATH) && ENVIRONMENT=testing $(GOTEST) -coverprofile=./tests/coverage-report.out ./...
	cd $(CODE_PATH) && $(GOCOVER) -func=./tests/coverage-report.out
	cd $(CODE_PATH) && $(GOCOVER) -html=./tests/coverage-report.out

.PHONY: e2e-test
e2e-test:
	@./local-scripts/integration-tests;\
	exit $$?;
