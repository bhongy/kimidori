# micro-frontend design notes

```sh
/_/service/{svc}/{feature}
/_/service/{svc}/{feature}/manifest.json
```

```json
// /_/service/{svc}/{feature}/manifest.json
{
  "name": "feature",
  "version": "1fc3...350f",
  "content": "/path/to/content", // HTML
  "css": {
    "asset.id": "/{svc}/static/4833a.css"
  },
  "js": {
    "asset.id": "/{svc}/static/a346c.js"
  }
}
```

```sh
curl "${svc}/path/to/content"

HTTP/1.1 200 OK
Content-Type: text/html
Upstream-Version: 1fc3...350f

<html>...
```
