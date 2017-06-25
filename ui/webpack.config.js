const webpack = require('webpack');
const path = require('path');
const ExtractTextPlugin = require("extract-text-webpack-plugin");

const nodeEnv = process.env.NODE_ENV || 'development';
const isProd = nodeEnv === 'production';

const common = {
  devtool: isProd ? 'hidden-source-map' : 'eval',
  context: __dirname,
  output: {
    path: path.join(__dirname, './build'),
    publicPath: '/assets/',
    filename: '[name].js',
  },

  devServer: {
    hot: true,
    host: "0.0.0.0",
    allowedHosts: [process.env.HOST],
    contentBase: './build/',
    publicPath: `https://${process.env.HOST}:8080/assets/`,
    historyApiFallback: true,
    stats: {
      assets: false,
      chunks: false,
      colors: false,
      modules: false, // hides JS and CSS info
    },
  },

  module: {
    rules: [
      {
        enforce: "pre",
        test: /\.js$/,
        exclude: /node_modules/,
        use: ["eslint-loader"],
      },
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
      config: {
        PORT: process.env.PORT || 3000,
      },
      'process.env': {
        NODE_ENV: JSON.stringify(nodeEnv),
      }
    }),
    new webpack.NoEmitOnErrorsPlugin(),
    new webpack.optimize.UglifyJsPlugin({
      compress: isProd,
      output: {
        comments: false
      },
      sourceMap: false
    }),
    new webpack.NamedModulesPlugin(),
  ]
};

if (isProd) {
  // Breaks HMR; only enable for smaller filesizes in prod
  common.plugins.push(new webpack.optimize.ModuleConcatenationPlugin());
}

module.exports = [
  Object.assign({}, common, {
    entry: {
      client: ['./src/entry.client.js'],
    },
  }),
  Object.assign({}, common, {
    entry: {
      server: ['./src/entry.server.js'],
    },
    target: 'node',
    node: {
      __dirname: false,
      __filename: false,
    },
    plugins: common.plugins.slice().concat([
      // debug < 2.6 references these which causes an error with SSR
      new webpack.DefinePlugin({
        window: {},
        navigator: false,
        document: false,
      }),
    ]),
  }),
];
