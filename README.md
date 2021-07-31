# kimidori

## Dev Notes

- group code by domain (service) not by type (database, website, etc)
  - each service owns its own data - never shares database
- top-level folders are domains (as in DDD domain)
- multiple services can be run locally in "prod-like" mode to provide an integrated environment when developing another service
  - the service being developed (dev mode) runs in vscode remote container
- the front-end and back-end layers are colocated in services
  - no front-end back-end coordinated deploy
  - we will not build a monolithic "webapp"

- unauth users (public) can do very little mainly to:
  - create an account
  - log in
  - recover password
  - ? view website/docs

- auth users have access to the app
