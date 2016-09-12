import { createStore, combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';
import { reducer } from 'tectonic';

const store = createStore(combineReducers({
  form: formReducer,
  tectonic: reducer,
}));

export default store;
