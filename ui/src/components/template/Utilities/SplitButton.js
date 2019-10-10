import React from 'react';
import classnames from 'classnames';

const sizeClasses = {
  small: 'btn-sm',
  large: 'btn-lg',
  normal: '',
};

export default function SplitButton({
  children,
  faIcon,
  variant,
  size,
  style,
  className,
  ...props
}) {
  const classNames = classnames(
    'btn',
    `btn-${variant}`,
    'btn-icon-split',
    sizeClasses[size],
    className
  );
  const styles = { cursor: 'pointer', ...style };
  return (
    <span className={classNames} style={styles} {...props}>
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
