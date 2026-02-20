# ========================================================================== #
# HELPERS
# ========================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


# ========================================================================== #
# DEVELOPMENT
# ========================================================================== #

## run/web: run the cmd/web application
.PHONY: run/web
run/web:
	go run ./cmd/web

## run/inject: run the cmd/examples/inject code
.PHONY: run/inject
run/inject: confirm
	go run cmd/examples/inject/main.go

## db/mysql: connect to the database usign mysql (need password)
.PHONY: db/mysql
db/mysql:
	mysql -D supermercado -u web -p

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${SUPERMERCADO_DB_DSN} up

## db/migrations/down: apply all down database migrations
.PHONY: db/migrations/down
db/migrations/down:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${SUPERMERCADO_DB_DSN} down


# ========================================================================== #
# QUALITY CONTROL
# ========================================================================== #

## fmt: format all .go files in the project directory
.PHONY: fmt
fmt:
	go fmt ./...

