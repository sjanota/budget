import React from 'react';
import { SplitButton } from './SplitButton';
import PropTypes from 'prop-types';
import { useDictionary } from './Lang';

export default function CancelButton({ onClick }) {
  const { buttons } = useDictionary();
  return (
    <SplitButton variant="danger" faIcon="times" size="small" onClick={onClick}>
      {buttons.cancel}
    </SplitButton>
  );
}

CancelButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};
