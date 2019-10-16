import React from 'react';
import { SplitButton } from './SplitButton';
import PropTypes from 'prop-types';

export default function CreateButton({ onClick }) {
  return (
    <SplitButton faIcon="plus" size="small" onClick={onClick}>
      Create
    </SplitButton>
  );
}

CreateButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};
