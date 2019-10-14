import Spinner from '../template/Utilities/Spinner';
import React from 'react';
import PropTypes from 'prop-types';

function ErrorMessageList({ errorMessage, subErrors }) {
  return (
    <>
      {errorMessage}
      <ul>
        {subErrors.map((e, idx) => (
          <li key={idx}>{e}</li>
        ))}
      </ul>
    </>
  );
}

function ErrorMessage({ error }) {
  const subErrors = error.networkError
    ? error.networkError.result.errors
    : error.graphQLErrors.map(e => `${e.path.join('.')}: ${e.message}`);
  return (
    <p className="text-danger">
      <i className="fas fa-fw fa-exclamation-triangle" />
      <ErrorMessageList errorMessage={error.message} subErrors={subErrors} />
    </p>
  );
}

export function WithQuery({ query, showError, children, ...props }) {
  const { loading, error } = query;
  return loading ? (
    <Spinner {...props} />
  ) : error ? (
    showError && <ErrorMessage error={error} />
  ) : (
    children(query)
  );
}

WithQuery.propTypes = {
  children: PropTypes.func.isRequired,
  query: PropTypes.shape({
    loading: PropTypes.bool.isRequired,
    error: PropTypes.any,
  }),
  showError: PropTypes.bool,
};

WithQuery.defaultProps = {
  showError: true,
};
