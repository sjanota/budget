import React from 'react';
import Card from 'react-bootstrap/Card';
import classnames from 'classnames';

export function Panel({ header, body, className }) {
  const classNames = classnames('card', 'shadow', 'mb-4', className);
  return (
    <div className={classNames}>
      <div className="card-header py-3">{header}</div>
      <Card.Body>{body}</Card.Body>
    </div>
  );
}

Panel.Title = function({ children }) {
  return <h6 className="m-0 font-weight-bold text-primary">{children}</h6>;
};
