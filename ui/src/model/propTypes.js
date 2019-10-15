import PropTypes from 'prop-types';

export const Amount = PropTypes.number;

export const Account = PropTypes.shape({
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
});

export const Expense = PropTypes.shape({
  id: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
  date: PropTypes.string,
  totalBalance: Amount.isRequired,
  location: PropTypes.string,
  account: Account,
});

export const Envelope = PropTypes.shape({
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
});

export const Category = PropTypes.shape({
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  envelope: Envelope,
});
