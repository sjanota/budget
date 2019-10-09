import React from 'react';
import { Button } from 'react-bootstrap';
import PropTypes from 'prop-types';

export default function TableButton({ faIcon, variant, ...props }) {
  return (
    <Button className="bg-transparent border-0 p-0 mx-1" {...props}>
      <i className={`fas fa-${faIcon} fa-fw text-${variant}`} />
    </Button>
  );
}

TableButton.propTypes = {
  faIcon: PropTypes.string.isRequired,
  variant: PropTypes.string.isRequired,
};
