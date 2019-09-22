import React from 'react';
import { Pencil } from '@primer/octicons-react';
import { OcticonButton } from './OcticonButton';
import PropTypes from 'prop-types';

export function EditButton({ onClick }) {
  return <OcticonButton icon={Pencil} action={'edit'} onClick={onClick} />;
}

EditButton.propTypes = {
  onClick: PropTypes.func,
};
