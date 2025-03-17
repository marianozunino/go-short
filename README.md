# go-short

A modern URL shortener service built with Go, inspired by bit.ly and similar services. Perfect for personal, self-hosted URL shortening needs.

## Features

- Shorten long URLs with one click
- Track click counts for each shortened URL
- Dark-themed, responsive UI with HTMX for interactivity
- Validates URL existence before shortening
- Efficiently stores and retrieves URLs with SQLite
- Uses MD5 hashing to prevent duplicate short URLs

## Dependencies

- [Echo](https://echo.labstack.com/) - High performance, minimalist Go web framework
- [templ](https://github.com/a-h/templ) - Typed templating language for Go
- [SQLite](https://www.sqlite.org/) - Self-contained, serverless database engine
- [sqlc](https://sqlc.dev/) - Generate type-safe Go from SQL
- [HTMX](https://htmx.org/) - HTML extensions for AJAX, WebSockets and more
- [Viper](https://github.com/spf13/viper) - Complete configuration solution

## Setup

1. Clone the repo:
   ```
   git clone https://github.com/marianozunino/go-short.git
   cd go-short
   ```

2. Get dependencies:
   ```
   go mod tidy
   ```

3. Run with task:
   ```
   # Setup and run development server
   task setup
   task dev

   # Or build and then run
   task build
   task serve
   ```

## Configuration

The application uses [Viper](https://github.com/spf13/viper) to manage configuration, supporting both JSON config files and environment variables.

### Environment Variables

| Environment Variable | Description                     | Default                    |
|----------------------|---------------------------------|----------------------------|
| PORT                 | HTTP server port                | 1323                       |
| DB_PATH              | Path to SQLite database         | "./db.sqlite"              |
| BASE_DOMAIN          | Base URL for shortened links    | "http://localhost:1323"    |

### Configuration File

You can also use a JSON configuration file to set these values.

## Usage

### Web Interface

1. Visit the homepage (default: http://localhost:1323)
2. Enter a valid URL in the input field
3. Click "Shorten"
4. Copy your shortened URL

### API Usage

```bash
# Shorten a URL (form submission)
curl -X POST -d "url=https://example.com/very/long/url/that/needs/shortening" http://localhost:1323/

# Access a shortened URL
curl -L http://localhost:1323/abc123
```

## Development

The project uses Task for managing development workflows:

```bash
# List all available tasks
task

# Generate templ templates
task templ

# Run database migrations
task migrate:up

# Reset database
task migrate:reset

# Generate Go code from SQL
task sqlc:generate
```

## License

MIT License - See [LICENSE](LICENSE) file for details
