# Design

Notes on design decisions.

# Host (App Shell)

The host page is responsible for:

- Routing i.e. loading apps based on the navigation
- Authentication and authorization
-

The host page never knows about the apps. As far as the host page concerns, the apps do not exist. Apps are deployed in dependently from each other and changing the apps does not require deploying the host. This means individual apps can go live, change, decommissioned in real time without needing to deploy the host.

> This design aims to prevent coordinated deploy through the host and allow the host to decouple from all the apps.

> Alternatively, we can "register services" at build time but it means the host now depends on all the apps (instead of the other way around) at build time and must be able to reach all of them in order to figure out the routes. In order words, we have a distributed monolith.

During the deployment, the apps will make an API call to register it's latest information about the route and how to load the "component" for the route to the Service Registry (back-end of the app host service).

There can be multiple "mounting" targets for an app (e.g. main, sidebar, etc).

Routes are namespaced by the apps. Since apps must have a unique name, the routes never collided. However, this might lead to apps relying on implementation detail (e.g. app names). We will see.

When loading an app, the host will pass the access token with the relevant scope to the app. The app will check with the authorization service (for every "access" request) before respond with the content. Authorization is the responsible of the apps.

## Why using iframe to render services

- protection for host environment (cookies, etc)
- event listeners without proper clean-ups from a service is isolated in its iframe

Problems to figure out:

- how (will) we allow an app to affect the host URL (e.g. in-app navigation with browser history)
- how to do cross-app navigation (how does app A knows what's the URL of app B is if it's resolved by the Service Registry) - e.g. "back" button or "home" button

Alternatively, the app shell can provide a `div#main` and load the entry .js of the client apps. The client app's entry bundle will `document.getElementById('main')` to mount the client app to the host. This approach is easier but exposes the host (app shell) environment to the client app.
