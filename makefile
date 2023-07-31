app = cook

.PHONY: direnv
direnv:
	direnv allow .

.PHONY: build
build:
	go build main.go

.PHONY: run
run:
	go run ./main.go --config-file=configs/config.yml

# OpenAPI
.PHONY: gen/openapihttp
gen/openapihttp:
	oapi-codegen -version
	oapi-codegen -generate types -o internal/infra/openapi/cooktypes.go -package openapi api/openapi/cook.yaml
	oapi-codegen -generate chi-server -o internal/infra/openapi/cookapi.go -package openapi api/openapi/cook.yaml
	oapi-codegen -generate types -o internal/client/openapi/cooktypes.go -package openapi api/openapi/cook.yaml
	oapi-codegen -generate client -o internal/client/openapi/cookapi.go -package openapi api/openapi/cook.yaml

.PHONY: sqlc/gen
gen/sqlc:
	sqlc generate -f ./configs/sqlc/pg.sqlc.yaml

.PHONY: pg/concat-migrations
pg/concat-migrations:
	# Merge all migrations into one single file and move it to `/.tmp`
	# Execute make command with sudo
	# psql: `\i /.tmp/pgall.sql`
	# WIP: Later a migrator will take care of this.
	cat /home/adrian/Projects/labs/foorester/cook/scripts/sql/pg/migrations/*.sql > pgall.sql
	mv pgall.sql /tmp

# API Calls
# CURL
.PHONY: api/create-book
api/create-book:
	./scripts/curl/create-book.sh -h localhost -p 8080 -n "Recipe Book One" -d "Favorite Recipes"

.PHONY: api/get-books
api/get-books:
	./scripts/curl/get-books.sh -h localhost -p 8080 -n "Recipe Book One" -d "Favorite Recipes"

# Testing
.PHONY: install/gomock
install/gomock:
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: mock/gen-ports
mock/gen-ports:
	# gomock is required (make install/gomock)
	mockgen -source=internal/domain/port/repo.go -destination=internal/infra/repo/pgx/repo_mock_test.go -package=pgx_test

# Docker
.PHONY: docker/dev
docker/dev:
	docker build . -t ak -f deployments/docker/dev/Dockerfile

.PHONY: docker/compose-up
docker/compose-up:
	# Run with sudo or, better, setup appropriate user permissions
	rm -rf tmp/postgres-data
	mkdir -p tmp/postgres-data
	chmod -R 777 tmp
	cp scripts/sql/docker/setup.sh tmp/sql
	cp scripts/sql/docker/2023062300-initial-setup.sql tmp/sql
	docker-compose -f ./deployments/docker/dev/docker-compose.yml up --build # --abort-on-container-exit --remove-orphans

.PHONY: docker/compose-down
docker/compose-down:
	docker-compose -f deployments/docker/dev/docker-compose.yml down

.PHONY: docker/reset-pg
docker/rest-pg:
	rm -rf tmp/postgres-data

.PHONY: docker/psql
docker/psql:
	sudo rm -rf ./tmp/postgres-data

# Testing
.PHONY: test
test:
	make -f makefile.test test-selected

.PHONY: docker/test
docker/test:
	make -f makefile.test compose-test

