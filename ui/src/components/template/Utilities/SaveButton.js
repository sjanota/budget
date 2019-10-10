import React from 'react';
import SplitButton from './SplitButton';
import PropTypes from 'prop-types';

export default function SaveButton(props) {
  return (
    <SplitButton faIcon="save" size="small" {...props}>
      Save
    </SplitButton>
  );
}

SaveButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};
