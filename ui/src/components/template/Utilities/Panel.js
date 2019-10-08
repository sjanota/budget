import React from 'react';
import Card from 'react-bootstrap/Card';

export default function Panel({ header, body }) {
  return (
    <div className="card shadow mb-4">
      <div className="card-header py-3">{header}</div>
      <Card.Body>{body}</Card.Body>
    </div>
  );
}

Panel.Title = function({ children }) {
  return <h6 className="m-0 font-weight-bold text-primary">{children}</h6>;
};
