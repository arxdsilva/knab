# knab

## Structure

```
knab/
├── bin
├── cmd
    ├── api
├── internal
    ├── domains
        ├── account
        ├── transaction
    ├── middleware
        ├── middlewares
        ├── router
    ├── handlers
        ├── rest
    ├── repository
    ├── storage
├── mocks
├── platform
    ├── postgres
```

Explaining the project structure:

1. bin/
    This is where the binaries are saved when using `go build` command.
2. cmd/
    Holds the main functions of the app, the instanciator for the app.
3. internal/
    Holds the `knab`'s project business logic and data flow. 
4. mocks/
    Centralizes all mocks necessary to test the app. This is optional and some developers prefer to have it in the mocked interface's package.
5. platform/
    Any tech that the project uses is here, the funcs written must be able to reuse across different pkgs.

## Adapters

Primary: everything that interacts with your API, eg: web client, cli, web servers
Secondary: everything that your API interacts with, eg: databases, other APIs, message queues

