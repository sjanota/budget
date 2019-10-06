import React, { useState } from 'react';
import classnames from 'classnames';
import { NavLink } from 'react-router-dom';

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
