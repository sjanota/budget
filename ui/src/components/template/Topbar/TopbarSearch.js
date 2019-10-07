import React from 'react';
import PropTypes from 'prop-types';

function TopbarSearch({ onSearch }) {
  return (
    <div className="input-group">
      <input
        type="text"
        className="form-control bg-light border-0 small"
        placeholder="Search for..."
        aria-label="Search"
        aria-describedby="basic-addon2"
      />
      <div className="input-group-append">
        <button className="btn btn-primary" type="button" onClick={onSearch}>
          <i className="fas fa-search fa-sm"></i>
        </button>
      </div>
    </div>
  );
}
TopbarSearch.propTypes = {
  onSearch: PropTypes.func,
};

function TopbarSearchExpanded(props) {
  return (
    <form className="d-none d-sm-inline-block form-inline mr-auto ml-md-3 my-2 my-md-0 mw-100 navbar-search">
      <TopbarSearch {...props} />
    </form>
  );
}

function TopbarSearchMinimized(props) {
  return (
    <li className="nav-item dropdown no-arrow d-sm-none">
      <a
        className="nav-link dropdown-toggle"
        href="#"
        id="searchDropdown"
        role="button"
        data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"
      >
        <i className="fas fa-search fa-fw"></i>
      </a>
      <div
        className="dropdown-menu dropdown-menu-right p-3 shadow animated--grow-in"
        aria-labelledby="searchDropdown"
      >
        <form className="form-inline mr-auto w-100 navbar-search">
          <TopbarSearch {...props} />
        </form>
      </div>
    </li>
  );
}

export default function ResponsiveTopbarSearch({ minified, ...props }) {
  return minified ? (
    <TopbarSearchMinimized {...props} />
  ) : (
    <TopbarSearchExpanded {...props} />
  );
}

ResponsiveTopbarSearch.propTypes = {
  minified: PropTypes.bool,
};
