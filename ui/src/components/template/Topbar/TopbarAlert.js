import React from 'react';
import classnames from 'classnames';
import PropTypes from 'prop-types';

export default function TopbarAlert({
  faIcon,
  variant,
  date,
  text,
  highlighted,
}) {
  const iconClasses = classnames('fas', 'text-white', faIcon);
  const iconBgClasses = classnames('icon-circle', `bg-${variant}`);
  const textClasses = classnames({ 'font-weight-bold': highlighted });
  return (
    <span className="topbar-alert dropdown-item d-flex align-items-center">
      <div className="mr-3">
        <div className={iconBgClasses}>
          <i className={iconClasses}></i>
        </div>
      </div>
      <div>
        <div className="small text-gray-500">{date}</div>
        <span className={textClasses}>{text}</span>
      </div>
    </span>
  );
}

TopbarAlert.propTypes = {
  date: PropTypes.string,
  faIcon: PropTypes.string,
  text: PropTypes.string,
  variant: PropTypes.string,
  highlighted: PropTypes.bool,
};

TopbarAlert.defaultProps = {
  variant: 'primary',
  highlighted: false,
};
