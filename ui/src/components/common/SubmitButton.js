import React from 'react';
import { Check } from '@primer/octicons-react';
import { OcticonButton } from './OcticonButton';
import PropTypes from 'prop-types';

export function SubmitButton({ onClick }) {
  return <OcticonButton icon={Check} action={'submit'} onClick={onClick} />;
}

SubmitButton.propTypes = {
  onClick: PropTypes.func,
};
