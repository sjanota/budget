import React, { useState, useContext } from 'react';
import PropTypes from 'prop-types';

const TemplateContext = React.createContext();

export function TemplateProvider({ children }) {
  const [sidebarToggled, setSidebarToggled] = useState(false);
  return (
    <TemplateContext.Provider
      value={{
        sidebarToggled,
        toggleSidebar: () => setSidebarToggled(current => !current),
      }}
    >
      {children}
    </TemplateContext.Provider>
  );
}

TemplateProvider.propTypes = {
  children: PropTypes.any,
};

export const useTemplate = () => useContext(TemplateContext);
