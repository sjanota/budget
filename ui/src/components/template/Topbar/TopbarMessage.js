import React from 'react';
import classnames from 'classnames';
import PropTypes from 'prop-types';

export default function TopbarMessage({
  highlighted,
  imgSrc,
  imgAlt,
  variant,
  text,
  author,
  when,
}) {
  const bgClass = {};
  bgClass[`bg-${variant}`] = variant;
  const indicatorClasses = classnames('status-indicator', bgClass);
  const textClasses = classnames({ 'font-weight-bold': highlighted });
  return (
    <a className="dropdown-item d-flex align-items-center" href="#">
      <div className="dropdown-list-image mr-3">
        <img className="rounded-circle" src={imgSrc} alt={imgAlt} />
        <div className={indicatorClasses}></div>
      </div>
      <div className={textClasses}>
        <div className="text-truncate">{text}</div>
        <div className="small text-gray-500">
          {author} · {when}
        </div>
      </div>
    </a>
  );
}

TopbarMessage.propTypes = {
  author: PropTypes.string,
  highlighted: PropTypes.bool,
  imgAlt: PropTypes.string,
  imgSrc: PropTypes.string,
  text: PropTypes.string,
  variant: PropTypes.string,
  when: PropTypes.string,
};

TopbarMessage.defaultProps = {
  highlighted: false,
};
