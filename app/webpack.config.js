'use strict';

const path = require('path');

module.exports = {
  entry: './src/index.tsx',
  output: {
    filename: '[name].js',
    chunkFilename: '[id].chunk.js',
    path: path.resolve(__dirname, 'dist/statics'),
    // webpack-dev-middleware uses this to filter requests for webpack assets
    // the idea is:
    //   - files emit to `dist/statics` in fs
    //   - request hit server for path `/statics/asset-name.ext`
    publicPath: '/statics',
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js'],
  },
  optimization: {
    runtimeChunk: true,
    splitChunks: {
      cacheGroups: {
        react: {
          test: /[\\/]node_modules[\\/](react|react-dom)[\\/]/,
          name: 'react',
          chunks: 'all'
        },
      },
    },
  },
};
