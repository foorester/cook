app = cook

.PHONY: direnv
direnv:
	direnv allow .

.PHONY: build
build:
	go build cmd/cook/main.go

.PHONY: run
run:
	go run ./cmd/$(app)/main.go --config-file=configs/config.yml

# OpenAPI
.PHONY: openapihttp
openapihttp:
	oapi-codegen -version
	oapi-codegen -generate types -o internal/infra/openapi/cooktypes.go -package openapi api/openapi/cook.yaml
	oapi-codegen -generate chi-server -o internal/infra/openapi/cookapi.go -package openapi api/openapi/cook.yaml
	oapi-codegen -generate types -o internal/client/openapi/cooktypes.go -package openapi api/openapi/cook.yaml
	oapi-codegen -generate client -o internal/client/openapi/cookapi.go -package openapi api/openapi/cook.yaml

.PHONY: gensqlcpg
gensqlcpg:
	sqlc generate -f ./configs/sqlc/pg.sqlc.yaml

.PHONY: pgall
pgall:
	# Merge all migrations into one single file and move it to `/.tmp`
	# Execute make command with sudo
	# psql: `\i /.tmp/pgall.sql`
	# WIP: Later a migrator will take care of this.
	cat /home/adrian/Projects/labs/foorester/cook/scripts/sql/pg/migrations/*.sql > pgall.sql
	mv pgall.sql /tmp

# API Call
# CURL
.PHONY: create-book
create-book:
	./scripts/curl/create-book.sh -h localhost -p 8080 -n "Recipe Book One" -d "Favorite Recipes"

# Testing
.PHONY: installgomock
installgomock:
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: genportmocks
genportmocks:
	# gomock is required (make installgomock)
	mockgen -source=internal/domain/port/repo.go -destination=internal/infra/repo/pgx/repo_mock_test.go -package=pgx_test

# Docker
.PHONY: docker-dev
docker-dev:
	docker build . -t ak -f deployments/docker/dev/Dockerfile

.PHONY: docker-compose
docker-compose:
	# Run with sudo or, better, setup appropriate user permissions
	rm -rf tmp/postgres-data
	mkdir -p tmp/postgres-data
	chmod -R 777 tmp
	cp scripts/sql/docker/setup.sh tmp/sql
	cp scripts/sql/docker/2023062300-initial-setup.sql tmp/sql
	docker-compose -f ./deployments/docker/dev/docker-compose.yml up --build # --abort-on-container-exit --remove-orphans

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose -f deployments/docker/dev/docker-compose.yml down

.PHONY: docker-reset-pg
docker-rest-pg:
	rm -rf tmp/postgres-data

.PHONY: docker-psql
docker-psql:
	sudo rm -rf ./tmp/postgres-data

# Testing
.PHONY: test
test:
	make -f makefile.test test-selected

.PHONY: docker-test
docker-test:
	make -f makefile.test compose-test
