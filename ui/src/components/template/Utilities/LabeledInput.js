import React from 'react';
import { Form } from 'react-bootstrap';
import PropTypes from 'prop-types';

export default function LabeledInput({ label, feedback, formData, ...props }) {
  return (
    <Form.Group className="mb-3">
      <Form.Label>{label}</Form.Label>
      <Form.Control ref={formData} defaultValue={formData.init} {...props} />
      <Form.Control.Feedback type="invalid">{feedback}</Form.Control.Feedback>
    </Form.Group>
  );
}

LabeledInput.propTypes = {
  formData: PropTypes.shape({ current: PropTypes.any, init: PropTypes.any }),
  feedback: PropTypes.string.isRequired,
  label: PropTypes.string.isRequired,
};
