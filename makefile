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
	oapi-codegen -generate types -o internal/infra/openapi/cooktypes.go -package openapi api/openapi/cook.yml
	oapi-codegen -generate chi-server -o internal/infra/openapi/cookapi.go -package openapi api/openapi/cook.yml
	oapi-codegen -generate types -o internal/infra/client/openapi/cooktypes.go -package openapi api/openapi/cook.yml
	oapi-codegen -generate client -o internal/infra/client/openapi/cookapi.go -package openapi api/openapi/cook.yml
