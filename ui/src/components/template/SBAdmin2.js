import React from 'react';
import { TemplateProvider } from './Context';
import PropTypes from 'prop-types';

export default function SBAdmin2({ children }) {
  return (
    <TemplateProvider>
      <div id="wrapper">{children}</div>
    </TemplateProvider>
  );
}

SBAdmin2.propTypes = {
  children: PropTypes.any,
};
