import Spinner from '../template/Utilities/Spinner';
import React from 'react';
import PropTypes from 'prop-types';

export default function WithQuery({ query, children, ...props }) {
  const { loading, error } = query;
  return loading ? (
    <Spinner {...props} />
  ) : error ? (
    <i className="fas fa-fw fa-exclamation-triangle text-danger" />
  ) : (
    children(query)
  );
}

WithQuery.propTypes = {
  children: PropTypes.node,
  query: PropTypes.shape({
    loading: PropTypes.bool.isRequired,
    error: PropTypes.any,
  }),
};
