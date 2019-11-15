import React from 'react';
import Card from 'react-bootstrap/Card';
import classnames from 'classnames';

export function Panel({
  header,
  headerClassName,
  body,
  bodyClassName,
  className,
}) {
  const classNames = classnames('card', 'shadow', 'mb-4', className);
  const headerClassNames = classnames(headerClassName, 'card-header');
  return (
    <div className={classNames}>
      {header && <div className={headerClassNames}>{header}</div>}
      {body && <Card.Body className={bodyClassName}>{body}</Card.Body>}
    </div>
  );
}

Panel.Title = function({ children }) {
  return <h6 className="m-0 font-weight-bold text-primary">{children}</h6>;
};

Panel.HeaderWithButton = function({ title, children }) {
  return (
    <div className="d-flex justify-content-between align-items-center">
      <Panel.Title>{title}</Panel.Title>
      <div>{children}</div>
    </div>
  );
};
