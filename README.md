# kimidori

## Dev Notes

- group code by domain (service) not by type (database, website, etc)
  - each service owns its own data - never shares database
- top-level folders are domains
- multiple services can be run locally in "prod-like" mode to provide an integrated environment when developing another service
  - the service being developed (dev mode) runs in vscode remote container

- unauth users (public) can do very little mainly to:
  - create an account
  - log in
  - recover password
  - ? view website/docs

- auth users have access to the app
