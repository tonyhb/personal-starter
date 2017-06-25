import React from 'react';
import ReactDOM from 'react-dom';
import App from './app.js';

const appEl = document.getElementById('app');

ReactDOM.render(<App />, appEl);

if (module.hot) {
  module.hot.accept('./app.js', () => {
    // If you use Webpack 2 in ES modules mode, you can
    // use <App /> here rather than require() a <NextApp />.

    ReactDOM.render(
      <App />,
      appEl,
    );
  });
}
