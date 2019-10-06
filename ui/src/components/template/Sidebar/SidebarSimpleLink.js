import React from 'react';
import classnames from 'classnames';
import { NavLink } from 'react-router-dom';
import PropTypes from 'prop-types';

export default function SidebarSimpleLink({ name, to, faIcon }) {
  const iconClasses = classnames('fas', 'fa-fw', faIcon);
  return (
    <li className="nav-item">
      <NavLink className="nav-link" exact to={to}>
        <i className={iconClasses} />
        <span>{name}</span>
      </NavLink>
    </li>
  );
}

SidebarSimpleLink.propTypes = {
  faIcon: PropTypes.string,
  name: PropTypes.string,
  to: PropTypes.string,
};
