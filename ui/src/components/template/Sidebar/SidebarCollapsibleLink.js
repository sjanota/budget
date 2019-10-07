import React from 'react';
import { NavLink } from 'react-router-dom';
import { useLocation } from 'react-router';
import classnames from 'classnames';
import PropTypes from 'prop-types';

function oneOfRoutsMatchesLocation(sections, location) {
  return sections.some(({ routes }) =>
    routes.some(({ to }) => location.pathname === to)
  );
}

export default function SidebarCollapsibleLink({
  name,
  parent,
  sections,
  faIcon,
}) {
  const location = useLocation();
  const isActive = oneOfRoutsMatchesLocation(sections, location);
  const classNames = classnames('nav-item', { active: isActive });
  const id = 'sidebar--' + name;
  const iconClasses = classnames('fas', 'fa-fw', `fa-${faIcon}`);
  return (
    <li className={classNames}>
      <span
        className="nav-link collapsed"
        data-toggle="collapse"
        data-target={'#' + id}
        aria-expanded="true"
        aria-controls={id}
      >
        <i className={iconClasses}></i>
        <span>{name}</span>
      </span>
      <div id={id} className="collapse" data-parent={'#' + parent}>
        <div className="bg-white py-2 collapse-inner rounded">
          {sections.map(({ name, routes }) => {
            return (
              <React.Fragment key={name}>
                <h6 className="collapse-header">{name}:</h6>
                {routes.map(({ to, label }) => {
                  return (
                    <NavLink key={label} className="collapse-item" to={to}>
                      {label}
                    </NavLink>
                  );
                })}
              </React.Fragment>
            );
          })}
        </div>
      </div>
    </li>
  );
}

SidebarCollapsibleLink.propTypes = {
  faIcon: PropTypes.string,
  name: PropTypes.string,
  parent: PropTypes.string,
  sections: PropTypes.arrayOf(
    PropTypes.shape({
      name: PropTypes.string,
      routes: PropTypes.arrayOf(
        PropTypes.shape({
          label: PropTypes.string,
          to: PropTypes.string,
        })
      ).isRequired,
    })
  ),
};
