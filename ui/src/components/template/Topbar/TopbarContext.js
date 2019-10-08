import React from 'react';
import PropTypes from 'prop-types';

function TopbarContextExpanded({ renderContext }) {
  return (
    <form className="d-none d-sm-inline-block form-inline mr-auto ml-md-3 my-2 my-md-0 mw-100 navbar-context">
      {renderContext()}
    </form>
  );
}

function TopbarContextMinimized({ faIcon, renderContext }) {
  return (
    <li className="nav-item dropdown no-arrow d-sm-none">
      <a
        className="nav-link dropdown-toggle"
        href="#"
        id="contextDropdown"
        role="button"
        data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"
      >
        <i className={`fas fa-${faIcon} fa-fw`}></i>
      </a>
      <div
        className="dropdown-menu dropdown-menu-right p-3 shadow animated--grow-in"
        aria-labelledby="contextDropdown"
      >
        <form className="form-inline mr-auto w-100 navbar-context">
          {renderContext()}
        </form>
      </div>
    </li>
  );
}

export default function TopbarContext({ minified, ...props }) {
  return minified ? (
    <TopbarContextMinimized {...props} />
  ) : (
    <TopbarContextExpanded {...props} />
  );
}

TopbarContext.propTypes = {
  minified: PropTypes.bool,
};
