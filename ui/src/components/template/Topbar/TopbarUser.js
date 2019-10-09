import React from 'react';

export default function TopbarUser() {
  return (
    <li className="nav-item dropdown no-arrow">
      <span
        className="nav-link dropdown-toggle"
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
          alt=""
          className="img-profile rounded-circle"
          src="https://source.unsplash.com/QAB-WJcbgJk/60x60"
        />
      </span>
      <div
        className="dropdown-menu dropdown-menu-right shadow animated--grow-in"
        aria-labelledby="userDropdown"
      >
        <span className="dropdown-item">
          <i className="fas fa-user fa-sm fa-fw mr-2 text-gray-400"></i>
          Profile
        </span>
        <span className="dropdown-item">
          <i className="fas fa-cogs fa-sm fa-fw mr-2 text-gray-400"></i>
          Settings
        </span>
        <span className="dropdown-item">
          <i className="fas fa-list fa-sm fa-fw mr-2 text-gray-400"></i>
          Activity Log
        </span>
        <div className="dropdown-divider"></div>
        <span
          className="dropdown-item"
          data-toggle="modal"
          data-target="#logoutModal"
        >
          <i className="fas fa-sign-out-alt fa-sm fa-fw mr-2 text-gray-400"></i>
          Logout
        </span>
      </div>
    </li>
  );
}
