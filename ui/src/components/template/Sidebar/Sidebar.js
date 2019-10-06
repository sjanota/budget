import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import classnames from 'classnames';
import PropTypes from 'prop-types';
import { SidebarGroup } from './SidebarGroup';

export default function Sidebar({ renderBrandName, renderBrandIcon, config }) {
  const [toggled, setToggled] = useState(false);
  const classNames = classnames(
    'navbar-nav',
    'bg-gradient-primary',
    'sidebar',
    'sidebar-dark',
    'accordion',
    { toggled }
  );
  return (
    <ul className={classNames} id="accordionSidebar">
      <Link
        className="sidebar-brand d-flex align-items-center justify-content-center"
        to="/"
      >
        <div className="sidebar-brand-icon rotate-n-15">
          {renderBrandIcon()}
        </div>
        <div className="sidebar-brand-text mx-3">{renderBrandName()}</div>
      </Link>

      <hr className="sidebar-divider my-0" />

      {config.map((group, idx) => (
        <SidebarGroup
          key={group.name || idx}
          group={group}
          isLast={idx !== config.length - 1}
        />
      ))}

      <hr className="sidebar-divider d-none d-md-block" />

      <div className="text-center d-none d-md-inline">
        <button
          className="rounded-circle border-0"
          id="sidebarToggle"
          onClick={() => setToggled(toggled => !toggled)}
        ></button>
      </div>
    </ul>
  );
}

Sidebar.propTypes = {
  config: PropTypes.arrayOf(SidebarGroup.propTypes.group).isRequired,
  renderBrandIcon: PropTypes.func.isRequired,
  renderBrandName: PropTypes.func.isRequired,
};
