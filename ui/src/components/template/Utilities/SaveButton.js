import React from 'react';
import SplitButton from './SplitButton';
import PropTypes from 'prop-types';

export default function SaveButton({ onClick }) {
  return (
    <SplitButton faIcon="save" size="small" onClick={onClick}>
      Save
    </SplitButton>
  );
}

SaveButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};
