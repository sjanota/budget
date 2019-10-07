import React from 'react';
import { useTemplate } from '../Context';
import './Topbar.css';

export default function Topbar({ renderMenus, renderContext, renderUser }) {
  const { toggleSidebar } = useTemplate();

  return (
    <nav className="navbar navbar-expand navbar-light bg-white topbar mb-4 static-top shadow">
      <button
        onClick={toggleSidebar}
        className="btn btn-link d-md-none rounded-circle mr-3"
      >
        <i className="fa fa-bars"></i>
      </button>

      {renderContext({ minified: false })}

      <ul className="navbar-nav ml-auto">
        {renderContext({ minified: true })}

        {renderMenus()}

        <div className="topbar-divider d-none d-sm-block" />

        {renderUser()}
      </ul>
    </nav>
  );
}
