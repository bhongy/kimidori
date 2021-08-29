const fs = require('fs');
const http = require('http');
const path = require('path');

const MIME = {
  '.html': 'text/html',
  '.css': 'text/css',
  '.js': 'text/javascript',
}

const root = path.join(__dirname, 'dist');
const server = http.createServer((req, res) => {
  fs.readFile(path.join(root, req.url), (err, d) => {
    if (err) {
      res.writeHead(404);
      res.end(JSON.stringify(err));
      return;
    }

    const ext = path.extname(req.url);
    const ct = MIME[ext];

    res.writeHead(200);
    res.setHeader('content-type', ct);
    res.end(d);
  });
});

server.listen(8000);
