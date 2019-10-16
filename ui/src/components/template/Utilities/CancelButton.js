import React from 'react';
import { SplitButton } from './SplitButton';
import PropTypes from 'prop-types';

export default function CancelButton({ onClick }) {
  return (
    <SplitButton variant="danger" faIcon="times" size="small" onClick={onClick}>
      Cancel
    </SplitButton>
  );
}

CancelButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};
