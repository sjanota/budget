import React from 'react';
import { SplitButton } from './SplitButton';
import PropTypes from 'prop-types';
import { useDictionary } from './Lang';

export default function SaveButton(props) {
  const { buttons } = useDictionary();

  return (
    <SplitButton faIcon="save" size="small" {...props}>
      {buttons.save}
    </SplitButton>
  );
}

SaveButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};
