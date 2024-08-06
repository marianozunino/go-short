# URL Shortener

This project is a URL shortener built in Go. It uses [templ.guide](https://templ.guide/) for HTML components, [sqlc](https://github.com/kyleconroy/sqlc) for generating type-safe SQL code, and SQLite as the database.

## Features

- Shorten long URLs
- Redirect to original URLs using shortened links
- Simple and clean HTML templates
- Lightweight and fast with Go and SQLite

## Requirements

- Go 1.16 or later
- [sqlc](https://github.com/kyleconroy/sqlc)
- [migrate](https://github.com/golang-migrate/migrate) for database migrations

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/url-shortener.git
    cd url-shortener
    ```

2. Install the dependencies:
    ```sh
    go mod download
    ```

3. Install `sqlc`:
    ```sh
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    ```

4. Install `migrate`:
    ```sh
    go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```

## Usage

### Running Migrations

Create a new migration:
```sh
migrate create -dir migrations -ext sql initial
```

Run the migrations:
```sh
migrate -database sqlite3://db.test -path migrations up
```

### Generating SQL Code with sqlc

To generate Go code from the SQL queries, run:
```sh
sqlc generate
```

### Running the Application

To start the URL shortener application, run:
```sh
go run main.go
```

## Configuration

The application can be configured using environment variables:

- `DATABASE_URL`: The URL for the SQLite database (default: `sqlite3://db.test`)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
