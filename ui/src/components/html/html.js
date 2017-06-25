import React from 'react';
import Helmet from 'react-helmet';

const HTML = ({ children }) => {
  const helmet = Helmet.renderStatic();
  const htmlAttrs = helmet.htmlAttributes.toComponent();
  const bodyAttrs = helmet.bodyAttributes.toComponent();

  return (
    <html {...htmlAttrs}>
      <head>
        {helmet.title.toComponent()}
        {helmet.meta.toComponent()}
        {helmet.link.toComponent()}
        <link rel='stylesheet' href='/assets/styles.css' />
      </head>
      <body {...bodyAttrs}>
        <div id="app" dangerouslySetInnerHTML={{ __html: children }} />
        <script src='/assets/client.js'></script>
      </body>
    </html>
  );
};

export default HTML;
