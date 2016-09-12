const webpack = require('webpack');
const path = require('path');
const ExtractTextPlugin = require("extract-text-webpack-plugin");

const nodeEnv = process.env.NODE_ENV || 'development';
const isProd = nodeEnv === 'production';

// postcss
const stylelint = require('stylelint');

module.exports = {
  devtool: isProd ? 'hidden-source-map' : 'cheap-eval-source-map',
  context: __dirname,
  entry: {
    app: ['./src/index.js'],
  },
  output: {
    path: path.join(__dirname, './build'),
    publicPath: '/assets/',
    filename: 'app.js'
  },
  module: {
    preLoaders: [
      // JS should be the first loader for dev-server.js
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        loaders: [
          'eslint-loader',
        ]
      },
    ],
    loaders: [
      // JS should be the first loader for dev-server.js
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        loaders: [
          'babel-loader',
        ]
      },
      {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract({
          fallbackLoader: 'style',
          loader: 'css?modules',
        })
      },
    ],
  },
  resolve: {
    extensions: ['', '.js', '.jsx'],
    modules: [
      path.resolve('./client'),
      'node_modules'
    ]
  },
  plugins: [
    new ExtractTextPlugin("styles.css"),
    new webpack.LoaderOptionsPlugin({
      minimize: true,
      debug: false,
    }),
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify(nodeEnv),
      }
    }),
    new webpack.NoErrorsPlugin(),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false
      },
      output: {
        comments: false
      },
      sourceMap: false
    }),
  ],
  devServer: {
    hot: true,
    contentBase: './build/',
    publicPath: '/assets/',
  },
  postcss: () => [
    stylelint(require('stylelint-config-standard')),
    require("postcss-reporter")({
      clearMessages: true,
      throwError: true,
    }),
  ],
};
