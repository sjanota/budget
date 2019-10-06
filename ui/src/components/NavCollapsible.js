import React from 'react';
import { NavLink } from 'react-router-dom';
import { useLocation } from 'react-router';
import classnames from 'classnames';

export default function NavCollapsible({
  id,
  name,
  labeledBy,
  parent,
  sections,
}) {
  const location = useLocation();
  let isActive = false;
  if (
    sections.some(({ routes }) =>
      routes.some(({ to }) => location.pathname === to)
    )
  ) {
    isActive = true;
  }

  console.log(location, isActive);

  const classNames = classnames('nav-item', { active: isActive });

  return (
    <li className={classNames}>
      <a
        className="nav-link collapsed"
        href="#"
        data-toggle="collapse"
        data-target={'#' + id}
        aria-expanded="true"
        aria-controls={id}
      >
        <i className="fas fa-fw fa-cog"></i>
        <span>{name}</span>
      </a>
      <div
        id={id}
        className="collapse"
        aria-labelledby={labeledBy}
        data-parent={'#' + parent}
      >
        {sections.map(({ name, routes }) => {
          return (
            <div key={name} className="bg-white py-2 collapse-inner rounded">
              <h6 className="collapse-header">{name}:</h6>
              {routes.map(({ to, label }) => {
                return (
                  <NavLink key={label} className="collapse-item" to={to}>
                    {label}
                  </NavLink>
                );
              })}
            </div>
          );
        })}
      </div>
    </li>
  );
}
