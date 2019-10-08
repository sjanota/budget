import React from 'react';
import { useTemplate } from '../Context';
import './Topbar.css';
import TopbarContext from './TopbarContext';

export default function Topbar({
  renderMenus,
  renderContext,
  renderUser,
  faIconContextMinified,
}) {
  const { toggleSidebar } = useTemplate();

  return (
    <nav className="navbar navbar-expand navbar-light bg-white topbar mb-4 static-top shadow">
      <button
        onClick={toggleSidebar}
        className="btn btn-link d-md-none rounded-circle mr-3"
      >
        <i className="fa fa-bars"></i>
      </button>

      <TopbarContext minified={false} renderContext={renderContext} />

      <ul className="navbar-nav ml-auto">
        <TopbarContext
          minified={true}
          renderContext={renderContext}
          faIcon={faIconContextMinified}
        />

        {renderMenus()}

        <div className="topbar-divider d-none d-sm-block" />

        {renderUser()}
      </ul>
    </nav>
  );
}
