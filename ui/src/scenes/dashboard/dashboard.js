import React from 'react';
import css from 'react-css-modules';

import styles from './dashboard.css';

console.log(styles);

const Dashboard = () => (
  <h1 styleName="base">Dashboard</h1>
);

export default css(styles)(Dashboard);
