import React from 'react';
import { Plus } from '@primer/octicons-react';
import { OcticonButton } from './OcticonButton';
import PropTypes from 'prop-types';

export function CreateButton({ onClick }) {
  return <OcticonButton icon={Plus} action={'create'} onClick={onClick} />;
}

CreateButton.propTypes = {
  onClick: PropTypes.func,
};
