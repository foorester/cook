# Cook - Recipe Manager

Cook is a simple example application that serves as a reference implementation for creating Go REST applications. It is designed to focus on microservices for CRUD operations providing a clear example of a RESTful microservice implemented in an elegant and straightforward way.

## Key Features

- Versionable REST API: Cook enables the implementation of a versioned REST API, allowing for future updates and enhancements.
- This application aims to reduce reliance on external dependencies while maintaining flexibility. It adopts established industry libraries for tasks such as configuration, routing, and interface definition language, such as OpenAPI and gRPC, when necessary. A more stringent implementation strategy, focusing on utilizing Go's standard library extensively can be found here [https://github.com/vanillazen/stl](https://github.com/vanillazen/stl).
- Use of Interfaces: Cook leverages interfaces to facilitate testing and enable the plugging in of alternative implementations, promoting modularity and flexibility.

## Extensibility and Code Generation

Cook is designed as a reference application that focuses on managing a couple of, recipes in this case, and provides a starting point for building Go based services.

Currently, the focus is on supporting SQLite, Postgres, and MongoDB. Since the database and repository functionalities are defined as interfaces in the service, there are no limitations preventing the creation of alternative implementations. 

There will be a separate code generator repository that utilizes Cook as a foundation for creating simple RESTful microservices specifically tailored for managing an intrinsically related group of resources (list-item, recipe-ingredient-directions, course-students, etc). The [code generator](https://github.com/foorester/crud) will provide developers with the ability to quickly generate the basic structure and functionality and extended later if required.

While the generated code will be optimized for managing a small set of resources, developers will be not limited to this constraints and will be able to adapt and modify the generated code to suit their specific and more complex use cases.

Finally, we recognize the significance of Test-driven development (TDD) principles. As the project's structure solidifies, our intention is to achieve comprehensive test coverage, addressing every aspect thoroughly.

## Usage
```shell
$ make run 
go run ./cmd/cook/main.go
[INF] 2023/06/23 12:01:16 cook starting...
[INF] 2023/06/23 12:01:16 user=cook password=cook dbname=foorester host=127.0.0.1 port=5432 search_path=cook
[INF] 2023/06/23 12:01:16 sqlc-db database connected!
[INF] 2023/06/23 12:01:16 cook started!
[INF] 2023/06/23 12:01:16 http-server started listening at localhost:8080

```

Make a `create-book` request in another terminal
```shell
$ make create-book 
./scripts/curl/create-book.sh -h localhost -p 8080 -n "Recipe Book One" -d "Favorite Recipes"
```

See the output
```shell
[DBG] 2023/06/23 12:01:51 ts: 2023/06/23 10:01:51, req-id: perun/uSr1bEITXA-000001, scheme: http, proto: HTTP/1.1, method: POST, addr: 127.0.0.1:37606, agent: curl/8.1.0, uri: http://localhost:8080/api/v1/books, status: 0, bytes: 0, elapsed: 0.470759ms
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
