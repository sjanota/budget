import React from 'react';

const sizeClasses = {
  small: 'btn-sm',
  large: 'btn-lg',
  normal: '',
};

export default function SplitButton({
  children,
  faIcon,
  variant,
  onClick,
  size,
}) {
  return (
    <span
      className={`btn btn-${variant} btn-icon-split ${sizeClasses[size]}`}
      style={{ cursor: 'pointer' }}
      onClick={onClick}
    >
      <span className="icon text-white-50">
        <i className={`fas fa-${faIcon}`}></i>
      </span>
      <span className="text">{children}</span>
    </span>
  );
}

SplitButton.defaultProps = {
  variant: 'primary',
  size: 'normal',
};
