const webpack = require('webpack');
const path = require('path');
const ExtractTextPlugin = require("extract-text-webpack-plugin");

const nodeEnv = process.env.NODE_ENV || 'development';
const isProd = nodeEnv === 'production';

module.exports = {
  devtool: isProd ? 'hidden-source-map' : 'eval',
  context: __dirname,
  entry: {
    app: ['./src/index.js'],
  },
  output: {
    path: path.join(__dirname, './build'),
    publicPath: '/assets/',
    filename: 'app.js'
  },

  devServer: {
    hot: true,
    host: "0.0.0.0",
    allowedHosts: [process.env.HOST],
    contentBase: './build/',
    publicPath: `https://${process.env.HOST}:8080/assets/`,
    historyApiFallback: true,
    stats: { chunks: false },
  },

  module: {
    rules: [
      // JS should be the first loader for dev-server.js
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: [
          'babel-loader',
        ]
      },
      {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: [
            {
              loader: 'css-loader',
              options: {
                importLoaders: 1,
                modules: true,
                localIdentName: '[local]', // injected into extracttext
              },
            },
            { loader: 'postcss-loader' },
          ]
        })
      },
    ],
  },
  resolve: {
    extensions: ['.js', '.jsx'],
    modules: [
      path.resolve('./src'),
      'node_modules'
    ]
  },
  plugins: [
    new ExtractTextPlugin({
      filename: 'styles.css',
      allChunks: true,
    }),
    new webpack.LoaderOptionsPlugin({
      minimize: true,
      debug: false,
    }),
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify(nodeEnv),
      }
    }),
    new webpack.NoEmitOnErrorsPlugin(),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false
      },
      output: {
        comments: false
      },
      sourceMap: false
    }),
    new webpack.NamedModulesPlugin(),
  ]
};
