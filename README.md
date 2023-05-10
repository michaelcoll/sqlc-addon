# sqlc-addon
A sqlc addon CLI that adds more code

## Supported Database
- SQLite

## What's done ?
This CLI will generate two new files :
- `connect.go` used to connect to the database
- `migration.go` used to initialize the database using the schema request that sqlc use

## How to install ?

```shell
go install github.com/michaelcoll/sqlc-addon@latest
```

## How to use ?
Place a new config file called `sqlc-addon.yaml` next to the sqlc config file.

The file should look like this
```yaml
version: "1"
addon_out: "internal/back/infrastructure/db"
database_name: "data.db"
```

Call the command in the same folder as the file `sqlc-addon.yaml` :
```shell
sqlc-addon generate
```
