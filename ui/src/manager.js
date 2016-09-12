import {
  Manager,
  BaseResolver,
} from 'tectonic';
import store from './store';

const manager = new Manager({
  drivers: {
  },
  resolver: new BaseResolver(),
  store,
});

export default manager;
