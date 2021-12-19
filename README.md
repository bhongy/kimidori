# kimidori

## Setting up local development cluster

All services in the project run in a kubernetes cluster. To develop locally, a local kubernetes cluster need to be setup

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

- Dependent services (:latest) run in the remote dev cluster shared by everyone but run one or a few services in development locally (or in your own remote dev instance).
  - services don't share state

- The state of the dev environment must always reflect the dev's git working directory state, which might or might not be committed to the central repository.
  - as an extension to above, when a dev pull from master, switch branch, rebase, the state of the dev environment must be updated accordingly i.e. two devs checking out the same commit with a clean git working directory will always have the same dev environment state

- service management UI
