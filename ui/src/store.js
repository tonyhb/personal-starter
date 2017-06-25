import { createStore, combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';
import { reducer } from 'tectonic';
import {
  Manager,
  BaseResolver,
} from 'tectonic';

const create = () => {
  const store = createStore(combineReducers({
    form: formReducer,
    tectonic: reducer,
  }));

  const manager = new Manager({
    drivers: {
    },
    resolver: new BaseResolver(),
    store,
  });

  return {
    store,
    manager,
  };
}

export default create;
