# Ghost Planet Backend

## Requirements

- Go 1.25.6 or newer
- PostgreSQL
- `psql` for the `make connect/db` helper

## Configuration

The API reads its PostgreSQL connection string from `GHOSTPLANET_DB_DSN` by default:

```sh
export GHOSTPLANET_DB_DSN='postgres://user:password@localhost:5432/ghostplanet?sslmode=disable'
```

Available runtime flags:

- `-port`: API port, default `4000`
- `-env`: environment label, default `development`
- `-db-dsn`: PostgreSQL DSN, default `$GHOSTPLANET_DB_DSN`
- `-db-max-open-conns`: default `25`
- `-db-max-idle-conns`: default `25`
- `-db-max-idle-time`: default `15m`

## Running Locally

```sh
go mod download
make run/api
```

The API listens on `http://localhost:4000` unless another port is provided:

```sh
go run ./cmd/api -port 4001
```

Connect to the configured database:

```sh
make connect/db
```
