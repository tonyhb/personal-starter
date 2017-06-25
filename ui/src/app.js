import React from 'react';
import { Provider } from 'react-redux';
import { BrowserRouter, Route } from 'react-router-dom';
import {
  Loader,
} from 'tectonic';

import store from './store';
import manager from './manager';

import Base from './scenes/base/base.js';
import Dashboard from './scenes/dashboard/dashboard.js';

const App = () => (
  <Provider store={ store }>
    <Loader manager={ manager }>
      <BrowserRouter>
        <Base>
          <Route exact pattern='/' component={ Dashboard } />
        </Base>
      </BrowserRouter>
    </Loader>
  </Provider>
);

export default App;
