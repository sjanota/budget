import React from 'react';
import classnames from 'classnames';
import PropTypes from 'prop-types';

function TopbarMenuCounter({ counter }) {
  const badge = counter >= 3 ? '3+' : '' + counter;
  return (
    counter > 0 && (
      <span className="badge badge-danger badge-counter">{badge}</span>
    )
  );
}

export function TopbarMenu({ name, faIcon, children, counter }) {
  const iconClasses = classnames('fas', 'fa-fw', `fa-${faIcon}`);
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
        <TopbarMenuCounter counter={counter} />
      </a>
      <div
        className="dropdown-list dropdown-menu dropdown-menu-right shadow animated--grow-in"
        aria-labelledby={id}
      >
        <h6 className="dropdown-header">{name}</h6>
        {children}
        <a className="dropdown-item text-center small text-gray-500" href="#">
          Show All
        </a>
      </div>
    </li>
  );
}

TopbarMenu.propTypes = {
  entries: PropTypes.arrayOf(
    PropTypes.shape({
      highlighted: PropTypes.bool,
      render: PropTypes.func.isRequired,
    })
  ),
  faIcon: PropTypes.string,
  name: PropTypes.string,
};
