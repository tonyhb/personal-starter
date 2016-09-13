const webpack = require('webpack');
const WebpackDevServer = require('webpack-dev-server');
const config = require('./webpack.config');

// Add webpack dev server HMR entrypoints
config.entry.app.unshift("webpack-dev-server/client?http://localhost:8080/", "webpack/hot/only-dev-server");
// add HMR
config.plugins.push(new webpack.HotModuleReplacementPlugin());
// And finally, add 'react-hot-loader/webpack' to JS loaders
config.module.loaders[0].loaders.push('react-hot-loader/webpack');

new WebpackDevServer(
  webpack(config),
  Object.assign({}, config.devServer)
).listen(8080);
