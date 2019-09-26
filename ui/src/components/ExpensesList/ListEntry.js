import React from 'react';
import * as MoneyAmount from '../../model/MoneyAmount';
import { Expense } from '../../model/propTypes';

export function ListEntry({ entry }) {
  return (
    <>
      <td>{entry.title}</td>
      <td>{entry.date}</td>
      <td>{MoneyAmount.format(entry.totalBalance)}</td>
      <td>{entry.location}</td>
      <td>{entry.account && entry.account.name}</td>
    </>
  );
}

ListEntry.propTypes = {
  entry: Expense.isRequired,
};
