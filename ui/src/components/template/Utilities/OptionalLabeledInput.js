import React, { useState } from 'react';
import { Form } from 'react-bootstrap';
import PropTypes from 'prop-types';

export default function OptionalLabeledInput({
  initEnabled,
  label,
  feedback,
  formData,
  ...props
}) {
  const [enabled, setEnabled] = useState(initEnabled);
  return (
    <Form.Group className="mb-3">
      <Form.Check custom type="switch">
        <Form.Check.Input checked={enabled} onChange={() => {}} />
        <Form.Check.Label
          onClick={() => {
            setEnabled(v => !v);
          }}
        >
          <Form.Label>{label}</Form.Label>
        </Form.Check.Label>
      </Form.Check>
      {enabled && (
        <>
          <Form.Control
            ref={formData}
            defaultValue={formData.init}
            {...props}
          />
          <Form.Control.Feedback type="invalid">
            {feedback}
          </Form.Control.Feedback>
        </>
      )}
    </Form.Group>
  );
}

OptionalLabeledInput.propTypes = {
  initEnabled: PropTypes.bool,
  formData: PropTypes.shape({ current: PropTypes.any, init: PropTypes.any }),
  feedback: PropTypes.string.isRequired,
  label: PropTypes.string.isRequired,
};
