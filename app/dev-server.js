'use strict';

const path = require('path');
const express = require('express');
const webpack = require('webpack');
const webpackConfig = require('./webpack.config.js');
const webpackDevMiddleware = require('webpack-dev-middleware');

const app = express();
const webpackCompiler = webpack({...webpackConfig, mode: 'development'});

app.use(
  webpackDevMiddleware(webpackCompiler, {
    stats: {
      colors: true,
    },
  })
);

app.get('*', (req, res) => {
  res.sendFile(path.resolve('index.html'));
});

const port = 8000;
app.listen(port, err => {
  if (err) {
    console.error(err);
  } else {
    console.log(`Dev server is running at: http://localhost:${port}`);
  }
});
