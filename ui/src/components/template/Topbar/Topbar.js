import React from 'react';
import { useTemplate } from '../Context';
import TopbarSearch from './TopbarSearch';
import classnames from 'classnames';
import PropTypes from 'prop-types';
import './Topbar.css';

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

function TopbarMenuCounter({ entries }) {
  const count = entries.filter(e => e.highlighted).length;
  const badge = count >= 3 ? '3+' : '' + count;
  return (
    count > 0 && (
      <span className="badge badge-danger badge-counter">{badge}</span>
    )
  );
}

function TopbarMenu({ name, faIcon, entries }) {
  const iconClasses = classnames('fas', 'fa-fw', faIcon);
  const id = `topbar--${name || faIcon}`;
  return (
    <li className="nav-item dropdown no-arrow mx-1">
      <a
        className="nav-link dropdown-toggle"
        href="#"
        id={id}
        role="button"
        data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"
      >
        <i className={iconClasses} />
        <TopbarMenuCounter entries={entries} />
      </a>
      <div
        className="dropdown-list dropdown-menu dropdown-menu-right shadow animated--grow-in"
        aria-labelledby={id}
      >
        <h6 className="dropdown-header">{name}</h6>
        {entries.map((e, idx) =>
          e.render({ highlighted: e.highlighted, key: idx })
        )}
        <a className="dropdown-item text-center small text-gray-500" href="#">
          Show All Alerts
        </a>
      </div>
    </li>
  );
}

function TopbarUser({ user }) {
  return (
    <li className="nav-item dropdown no-arrow">
      <a
        className="nav-link dropdown-toggle"
        href="#"
        id="userDropdown"
        role="button"
        data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"
      >
        <span className="mr-2 d-none d-lg-inline text-gray-600 small">
          Valerie Luna
        </span>
        <img
          className="img-profile rounded-circle"
          src="https://source.unsplash.com/QAB-WJcbgJk/60x60"
        />
      </a>
      <div
        className="dropdown-menu dropdown-menu-right shadow animated--grow-in"
        aria-labelledby="userDropdown"
      >
        <a className="dropdown-item" href="#">
          <i className="fas fa-user fa-sm fa-fw mr-2 text-gray-400"></i>
          Profile
        </a>
        <a className="dropdown-item" href="#">
          <i className="fas fa-cogs fa-sm fa-fw mr-2 text-gray-400"></i>
          Settings
        </a>
        <a className="dropdown-item" href="#">
          <i className="fas fa-list fa-sm fa-fw mr-2 text-gray-400"></i>
          Activity Log
        </a>
        <div className="dropdown-divider"></div>
        <a
          className="dropdown-item"
          href="#"
          data-toggle="modal"
          data-target="#logoutModal"
        >
          <i className="fas fa-sign-out-alt fa-sm fa-fw mr-2 text-gray-400"></i>
          Logout
        </a>
      </div>
    </li>
  );
}

export default function Topbar({ onSearch, config, user }) {
  const { toggleSidebar } = useTemplate();

  return (
    <nav className="navbar navbar-expand navbar-light bg-white topbar mb-4 static-top shadow">
      <button
        onClick={toggleSidebar}
        className="btn btn-link d-md-none rounded-circle mr-3"
      >
        <i className="fa fa-bars"></i>
      </button>

      <TopbarSearchExpanded onSearch={onSearch} />

      <ul className="navbar-nav ml-auto">
        <TopbarSearchMinimized onSearch={onSearch} />

        {config.map((group, idx) => (
          <TopbarMenu key={group.name || idx} {...group} />
        ))}

        <div className="topbar-divider d-none d-sm-block" />

        <TopbarUser user={user} />
      </ul>
    </nav>
  );
}
