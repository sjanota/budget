import React from 'react';

export default function PageHeaderWithButton({
  title,
  btnText,
  faIcon,
  variant,
  onClick,
}) {
  return (
    <div className="d-sm-flex align-items-center justify-content-between mb-4">
      <h1 className="h3 mb-4 text-gray-800">{title}</h1>
      <span
        style={{ cursor: 'pointer' }}
        onClick={onClick}
        className={`d-none d-sm-inline-block btn btn-sm btn-${variant} shadow-sm`}
      >
        <i className={`fas ${faIcon} fa-sm text-white-50`} /> {btnText}
      </span>
    </div>
  );
}

PageHeaderWithButton.defaultProps = {
  variant: 'primary',
};
