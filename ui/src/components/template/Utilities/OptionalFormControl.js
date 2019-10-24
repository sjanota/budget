import React, { useState } from 'react';
import { Form, Row } from 'react-bootstrap';
import PropTypes from 'prop-types';
import { FormControl } from './FormControl';

export function OptionalFormControl({
  initEnabled,
  label,
  inline,
  formData,
  ...props
}) {
  const [enabled, setEnabled] = useState(initEnabled);
  const toggleEnabled = () => setEnabled(v => !v);
  return (
    <Form.Group className="mb-3" as={!!inline && Row}>
      <Form.Label column={!!inline}>
        <Form.Check custom type="switch">
          <Form.Check.Input checked={enabled} onChange={toggleEnabled} />
          <Form.Check.Label onClick={toggleEnabled}>{label}</Form.Check.Label>
        </Form.Check>
      </Form.Label>
      {enabled && (
        <FormControl.Input
          autoFocus
          formData={formData}
          inline={inline}
          {...props}
        />
      )}
    </Form.Group>
  );
}

OptionalFormControl.propTypes = {
  initEnabled: PropTypes.bool,
  formData: PropTypes.shape({ current: PropTypes.any, init: PropTypes.any }),
  feedback: PropTypes.string.isRequired,
  label: PropTypes.string.isRequired,
  inline: PropTypes.number,
};

OptionalFormControl.defaultProps = {
  inline: 0,
};
