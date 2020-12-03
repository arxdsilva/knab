# knab

## Structure

```
knab/
├── cmd
    ├── api
├── internal
    ├── domains
        ├── account
        ├── transaction
    ├── handlers
    ├── middlewares
    ├── repository
├── mocks
├── platform
    ├── config
    ├── mocks
├── queries
├── migrations
```

project structure:

1. cmd/
    Holds the main function of the app, the instanciator for the app.
2. internal/
    Holds the `knab`'s project business logic and data flow. 
3. queries/ and migrations/
    Prest's folders that allows custom queries and db migrations
4. platform/
    Code related to project, but able to reuse across different pkgs. Mocks are centralized here but could be distributed into their interface's pkgs.

## Packages and Golang version

- Go used was version `1.15.2`
- Main packages include:
    - `prest` for db migrations, route logs, allows many other fast dev features such as custom queries
    - `kpango/glg` logging porpuses
    - `stretchr/testify` for testing assertions
- Database used is `postgresql` on version 13

## Testing

The API can be tested by using `go test -v ./...` on root folder, also there's a `insomnia.json` file that allows manual tests.

## Running

1. Copy the config file `config.toml.sample` to `config.toml` using `cp config.toml.sample config.toml`;
2. Run the postgres container with `make postgres`;
3. Run the migrations using `make migrate`;
4. Run the app with `make run` and it'll be up on port `8888`;
