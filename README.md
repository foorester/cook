# Cook - Recipe Manager

Cook is a simple example application that serves as a reference implementation for creating Go REST applications. It is designed to focus on microservices for CRUD operations providing a clear example of a RESTful microservice implemented in an elegant and straightforward way.

## Key Features

- Versionable REST API: Cook enables the implementation of a versioned REST API, allowing for future updates and enhancements.
- This application aims to reduce reliance on external dependencies while maintaining flexibility. It adopts established industry libraries for tasks such as configuration, routing, and interface definition language, such as OpenAPI and gRPC, when necessary. A more stringent implementation strategy, focusing on utilizing Go's standard library extensively can be found here [https://github.com/vanillazen/stl](https://github.com/vanillazen/stl).
- Use of Interfaces: Cook leverages interfaces to facilitate testing and enable the plugging in of alternative implementations, promoting modularity and flexibility.

## Extensibility

Cook is designed as a reference application that focuses on managing recipes and provides a starting reference point for building Go based services.

Currently, the focus is on supporting SQLite, Postgres, and MongoDB. Since the database and repository functionalities are defined as interfaces in the service, there are no limitations preventing the creation of alternative implementations. 

Our service includes a package that follows ROCA recommendations and implements a client-side architecture with a traditional HTML Server-Side Rendering (SSR) approach. For enhanced dynamism, this implementation may utilize popular libraries such as React, Vue, or Svelte to provide individual pages with interactive elements. The goal is to create a structured and efficient web application experience, similar to platforms like GitHub, where resource-oriented principles are combined with dynamic components for improved performance, scalability, and user engagement.

The ability to validate JWT tokens is also planned.

Finally, we recognize the significance of Test-driven development (TDD) principles. As the project's structure solidifies, our intention is to achieve comprehensive test coverage, addressing every aspect thoroughly.

## Usage
```shell
$ make run 
go run ./cmd/cook/main.go
go run ./main.go --config-file=configs/config.yml
[INF] 2023/07/30 19:24:51 cook starting...
[INF] 2023/07/30 19:24:51 migrator started
[INF] 2023/07/30 19:24:51 migrator database connected
[INF] 2023/07/30 19:24:51 Migration 'users' already applied
[INF] 2023/07/30 19:24:51 Migration 'books' already applied
[INF] 2023/07/30 19:24:51 Migration 'recipes' already applied
[INF] 2023/07/30 19:24:51 Migration 'ingredients' already applied
[INF] 2023/07/30 19:24:51 Migration 'steps' already applied
[INF] 2023/07/30 19:24:51 seeder started
[INF] 2023/07/30 19:24:51 seeder database connected
[INF] 2023/07/30 19:24:51 Seed 'users' already applied
[DBG] 2023/07/30 19:24:51 user=cook password=cook dbname=foorester host=localhost port=5432 search_path=cook
[INF] 2023/07/30 19:24:51 sqlc-db database connected!
[INF] 2023/07/30 19:24:51 cook started!
[INF] 2023/07/30 19:24:51 http-server started listening at localhost:8080
```

Make a `create-book` request in another terminal
```shell
$ make create-book 
./scripts/curl/create-book.sh -h localhost -p 8080 -n "Recipe Book One" -d "Favorite Recipes"
```

See the output
```shell
[DBG] 2023/07/30 19:25:08 ts: 2023/06/23 10:01:51, req-id: perun/uSr1bEITXA-000001, scheme: http, proto: HTTP/1.1, method: POST, addr: 127.0.0.1:37606, agent: curl/8.1.0, uri: http://localhost:8080/api/v1/books, status: 0, bytes: 0, elapsed: 0.470759ms
```


## Highlighted files
### HTTP handlers
* [http/handler.go](internal/infra/http/handler.go)

### Business logic
* [service/recipe.go](internal/core/service/recipe.go)

## Notes
This projects utilizes a customized fork of the OpenAPI generator for its server and client interface needs. While the original version of the generator, available at https://github.com/deepmap/oapi-codegen, remains a viable option, we have opted to use [this forked version](https://github.com/foorester/oapi-codegen) for improved code clarity.

One concrete example highlighting is the transformation of a server interface function originally named `DeleteRecipeBooksBookIdRecipesRecipeIdIngredientsIngredientId` into `DeleteIngredient`.

## License

This project is licensed under the MIT License. Feel free to use and modify it as per the terms of the license.
