# authentication

authentication answers whether a credential is valid i.e. the user providing the credential is logged in. It does not provide information about the user.

It is in the foundation service layer and never to be hit by a public traffic (i.e. hit by webapp).

## Before using devcontainer for the first time

- make a copy of `.env.example` to `.env` and fill in the appropriate environment variables

## Running database migration

```sh
$ cd internal/data/db/migrations

# applies all migrations
$ tern migrate

# revert all migrations
$ tern migrate -d 0

# see current migration status
$ tern status

# applies all migrations of the test db
$ tern migrate -c tern-testdb.conf

# see current migration status of the test db
$ tern status -c tern-testdb.conf
```
