import React from 'react';
import PropTypes from 'prop-types';
import { Button } from 'react-bootstrap';
import classnames from 'classnames';

export default function RefreshButton({ className, ...props }) {
  const classes = classnames('btn-sm', 'btn-secondary', className);
  return (
    <Button className={classes} {...props}>
      <i className="fas fa-fw fa-sync-alt" />
    </Button>
  );
}

RefreshButton.propTypes = {
  className: PropTypes.string,
};
