# authentication

authentication answers whether a credential is valid i.e. the user providing the credential is logged in. It does not provide information about the user.

## Before using devcontainer for the first time

- make a copy of `.env.example` to `.env` and fill in the appropriate environment variables

## Running database migration

> ❗ the migration commands must be run (cwd) from root kimidori/authentication

```sh
# applies all migrations
$ go run ./internal/data/migrations up

# revert all migrations
$ go run ./internal/data/migrations down

# see all commands
$ go run ./internal/data/migrations
```
