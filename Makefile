GOOSE_DRIVER:="mysql"
GOOSE_DB_STRING:='root:root@tcp(db:3306)/711test'
up:
	docker-compose up -d
down:
	docker-compose down
up-no-cache: down
	docker-compose build --no-cache && \
	docker-compose up -d

build:
	docker-compose build
bash:
	docker-compose exec go bash

mg-up:
	@goose -dir ./migrations ${GOOSE_DRIVER} ${GOOSE_DB_STRING} up

mg-create:
	@goose -dir ./migrations create ${name} sql

mg-redo:
	@goose -dir ./migrations $(GOOSE_DRIVER) ${GOOSE_DB_STRING} redo

mg-down:
	@goose -dir ./migrations $(GOOSE_DRIVER) ${GOOSE_DB_STRING} down

mg-status:
	@goose -dir ./migrations $(GOOSE_DRIVER) ${GOOSE_DB_STRING} status	

test1:
	@echo ${GOOSE_DB_STRING}