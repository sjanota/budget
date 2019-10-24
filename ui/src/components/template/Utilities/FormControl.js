import React from 'react';
import { Form, Row, Col } from 'react-bootstrap';
import PropTypes from 'prop-types';

export function FormControl({ label, inline, ...props }) {
  return (
    <Form.Group className="mb-3" as={!!inline && Row}>
      <Form.Label column={!!inline}>{label}</Form.Label>
      <FormControl.Input inline={inline} {...props} />
    </Form.Group>
  );
}

FormControl.Input = ({ inline, formData, feedback, ...props }) => {
  const wrap = inline ? c => <Col sm={inline}>{c}</Col> : c => c;

  return wrap(
    <>
      <Form.Control
        ref={formData}
        defaultValue={formData.default()}
        {...props}
      />
      <Form.Control.Feedback type="invalid">{feedback}</Form.Control.Feedback>
    </>
  );
};

FormControl.propTypes = {
  formData: PropTypes.shape({ current: PropTypes.any, init: PropTypes.any }),
  feedback: PropTypes.string.isRequired,
  label: PropTypes.string.isRequired,
  inline: PropTypes.number,
};

FormControl.defaultProps = {
  inline: 0,
};
