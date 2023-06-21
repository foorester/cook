app = cook

.PHONY: direnv
direnv:
	direnv allow .

.PHONY: build
build:
	go build ./...

.PHONY: run
run:
	go run ./cmd/$(app)/main.go

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
	# Merge all migrations into one single file and move it to `/tmp`
	# Execute make command with sudo
	# psql: `\i /tmp/pgall.sql`
	cat /home/adrian/Projects/labs/foorester/cook/scripts/sql/pg/migrations/*.sql > pgall.sql
	mv pgall.sql /tmp

# CURL
.PHONY: create-book
create-book:
	./scripts/curl/create_book.sh -h localhost -p 8080 -n "Recipe Book One" -d "Favorite Recipes"
