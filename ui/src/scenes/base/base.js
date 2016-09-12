import React from 'react';

const Base = ({ children }) => (
  <div>
    <h1>LOL</h1>
    { children }
  </div>
);

Base.propTypes = {
  children: React.PropTypes.node,
};

export default Base;
