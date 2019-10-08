import React from 'react';
import classnames from 'classnames';

export default function Spinner({ size, variant }) {
  const classNames = classnames('spinner-border', {
    [`spinner-border-${size}`]: size,
    [`text-${variant}`]: variant,
  });
  return (
    <div className={classNames} role="status">
      <span className="sr-only">Loading...</span>
    </div>
  );
}
