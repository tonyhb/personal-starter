import React from 'react';
import css from 'react-css-modules';
import { Link } from 'react-router-dom';
import styles from './base.css';

const Base = ({ children }) => (
  <div>
    <div styleName="styles.header">
      <Link to='/' styleName='styles.dashboard'>Dashboard</Link>
    </div>
    { children }
  </div>
);

export default Base;
