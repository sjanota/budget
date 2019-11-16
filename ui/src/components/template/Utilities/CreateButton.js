import React from 'react';
import { SplitButton } from './SplitButton';
import PropTypes from 'prop-types';
import { useDictionary } from './Lang';

export default function CreateButton({ onClick }) {
  const { buttons } = useDictionary();
  return (
    <SplitButton faIcon="plus" size="small" onClick={onClick}>
      {buttons.create}
    </SplitButton>
  );
}

CreateButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};
