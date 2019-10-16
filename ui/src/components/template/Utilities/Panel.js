import React from 'react';
import Card from 'react-bootstrap/Card';
import classnames from 'classnames';

export function Panel({ header, body, className }) {
  const classNames = classnames('card', 'shadow', 'mb-4', className);
  return (
    <div className={classNames}>
      {header && <div className="card-header py-3">{header}</div>}
      {body && <Card.Body>{body}</Card.Body>}
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
