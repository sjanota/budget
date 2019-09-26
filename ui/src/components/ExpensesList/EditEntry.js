import React from 'react';
import PropTypes from 'prop-types';
import { Expense } from '../../model/propTypes';
import * as MoneyAmount from '../../model/MoneyAmount';

export function EditEntry({ entry, setEntry }) {
  function setValue(value) {
    return setEntry(e => ({ ...e, ...value }));
  }
  return (
    <>
      <td>
        <input
          value={entry.title}
          onChange={event => setValue({ title: event.target.value })}
          type={'text'}
        />
      </td>
      <td>
        <input
          value={entry.date || ''}
          onChange={event => setValue({ date: event.target.value })}
          type={'date'}
        />
      </td>
      <td>
        <input
          value={MoneyAmount.format(entry.totalBalance)}
          onChange={event =>
            setValue({ totalBalance: MoneyAmount.parse(event.target.value) })
          }
          type={'number'}
        />
      </td>
      <td>
        <input
          value={entry.location || ''}
          onChange={event => setValue({ location: event.target.value })}
          type={'text'}
        />
      </td>
      <td />
    </>
  );
}

EditEntry.propTypes = {
  entry: Expense.isRequired,
  setEntry: PropTypes.func.isRequired,
};
