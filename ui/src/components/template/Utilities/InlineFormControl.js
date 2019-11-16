import React from 'react';
import { Form, Row, Col } from 'react-bootstrap';
import PropTypes from 'prop-types';

export function InlineFormControl({ label, size, feedback, children }) {
  return (
    <Form.Group as={Row}>
      <Form.Label column>{label}</Form.Label>
      <Col sm={size}>
        {children}
        {feedback && (
          <Form.Control.Feedback type="invalid">
            {feedback}
          </Form.Control.Feedback>
        )}
      </Col>
    </Form.Group>
  );
}

InlineFormControl.propTypes = {
  children: PropTypes.element,
  feedback: PropTypes.string,
  label: PropTypes.string.isRequired,
  size: PropTypes.number,
};

InlineFormControl.defaultProps = {
  inline: 0,
};
