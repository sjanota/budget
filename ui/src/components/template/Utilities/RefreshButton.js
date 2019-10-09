import React from 'react';
import PropTypes from 'prop-types';
import { Button } from 'react-bootstrap';

export default function RefreshButton(props) {
  return (
    <Button className="btn-sm btn-secondary" {...props}>
      <i className="fas fa-fw fa-sync-alt" />
    </Button>
  );
}

RefreshButton.propTypes = {
  onClick: PropTypes.func.isRequired,
};
