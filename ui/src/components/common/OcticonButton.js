import React from 'react';
import Octicon from '@primer/octicons-react';
import { Button } from 'react-bootstrap';
import PropTypes from 'prop-types';

const SMALL = {
  bootstrap: 'sm',
  octicon: 'small',
};

const Size = PropTypes.shape({
  bootstrap: PropTypes.string,
  octicon: PropTypes.string,
});

export function OcticonButton({ icon, action, onClick, size }) {
  return (
    <Button
      size={size.bootstrap}
      variant={'link'}
      data-action={action}
      onClick={onClick}
    >
      <Octicon icon={icon} size={size.octicon} />
    </Button>
  );
}

OcticonButton.propTypes = {
  icon: PropTypes.any.isRequired,
  action: PropTypes.string,
  onClick: PropTypes.func,
  size: Size,
};

OcticonButton.defaultProps = {
  size: SMALL,
};
