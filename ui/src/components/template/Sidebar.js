import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import SidebarCollapsibleLink from './SidebarCollapsibleLink';
import classnames from 'classnames';
import SidebarSimpleLink from './SidebarSimpleLink';

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
        <>
          {group.name && <div className="sidebar-heading">{group.name}</div>}
          {group.entries.map(entry =>
            entry.to !== undefined ? (
              <SidebarSimpleLink key={entry.name} {...entry} />
            ) : (
              <SidebarCollapsibleLink
                key={entry.name}
                parent="accordionSidebar"
                {...entry}
              />
            )
          )}
          {idx !== config.length - 1 && <hr className="sidebar-divider" />}
        </>
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
