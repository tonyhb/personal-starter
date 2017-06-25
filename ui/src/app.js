import React from 'react';
import { Provider } from 'react-redux';
import { BrowserRouter, Route } from 'react-router-dom';
import { Loader } from 'tectonic';

import Base from './scenes/base/base.js';
import Dashboard from './scenes/dashboard/dashboard.js';

const App = ({ store, manager }) => (
  <Provider store={ store }>
    <Loader manager={ manager }>
      <Base>
        <Route exact pattern='/' component={ Dashboard } />
      </Base>
    </Loader>
  </Provider>
);

export default App;
