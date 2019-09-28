import PropTypes from 'prop-types';

export const MoneyAmount = PropTypes.shape({
  integer: PropTypes.number.isRequired,
  decimal: PropTypes.number.isRequired,
});

export const Account = PropTypes.shape({
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
});

export const Expense = PropTypes.shape({
  id: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
  date: PropTypes.string,
  totalBalance: MoneyAmount.isRequired,
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
