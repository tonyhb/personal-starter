import React from 'react';
import { Provider } from 'react-redux';
import { Router, Route, browserHistory } from 'react-router';
import {
  Loader,
} from 'tectonic';
import 'sanitize.css/sanitize.css';

import store from './store';
import manager from './manager';

import Base from './scenes/base/base.js';
import Dashboard from './scenes/dashboard/dashboard.js';

const app = (
  <Provider store={ store }>
    <Loader manager={ manager }>
      <Router history={ browserHistory }>
        <Route component={ Base }>
          <Route path="/" component={ Dashboard } />
        </Route>
      </Router>
    </Loader>
  </Provider>
);

export default app;
