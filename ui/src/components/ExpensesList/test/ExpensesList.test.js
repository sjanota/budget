import { mount } from 'enzyme';
import React from 'react';
import { MockedProvider } from '@apollo/react-testing';
import {
  expense1,
  expense2,
  mockDeleteExpense,
  mockExpensesEvent,
  mockQueryExpenses,
} from './ExpensesList.test.mocks';
import { updateComponent } from '../../../testing';
import { createMockLink } from '../../../testing/apollo';
import ExpensesList from '../ExpensesList';
import { MockedBudgetProvider } from '../../context/budget/budget.testing';

it('displays loading initially', async () => {
  const { link } = createMockLink([]);
  console.error = jest.fn();
  const component = mount(
    <MockedBudgetProvider>
      <MockedProvider link={link}>
        <ExpensesList />
      </MockedProvider>
    </MockedBudgetProvider>
  );

  expect(component.find('tbody tr')).toHaveLength(0);
  expect(component.find('p')).toExist();
  expect(component.find('p')).toHaveText('Loading...');
});

it('displays error if occurs', async () => {
  const { link } = createMockLink([]);
  console.error = jest.fn();
  const component = mount(
    <MockedBudgetProvider>
      <MockedProvider link={link}>
        <ExpensesList />
      </MockedProvider>
    </MockedBudgetProvider>
  );
  await updateComponent(component);

  expect(console.error).toHaveBeenCalled();
  expect(component.find('tbody tr')).toHaveLength(0);
  expect(component.find('p')).toExist();
  expect(component.find('p')).toHaveText('Error :(');
});

describe('When data is loaded', () => {
  const deleteMock = mockDeleteExpense(expense1.id);
  const { link, sendEvent } = createMockLink([
    mockQueryExpenses([expense1]),
    deleteMock,
  ]);
  const component = mount(
    <MockedBudgetProvider>
      <MockedProvider link={link}>
        <ExpensesList />
      </MockedProvider>
    </MockedBudgetProvider>
  );

  it('displays queried data', async () => {
    await updateComponent(component);
    expect(component.find('tbody tr')).toHaveLength(1);
  });

  it('updates list on CREATED', async () => {
    sendEvent(mockExpensesEvent({ type: 'CREATED', expense: expense2 }));
    await updateComponent(component);

    expect(component.find('tbody tr')).toHaveLength(2);
  });

  it('updates list on DELETED', async () => {
    sendEvent(mockExpensesEvent({ type: 'DELETED', expense: expense2 }));
    await updateComponent(component);

    expect(component.find('tbody tr')).toHaveLength(1);
  });

  it('triggers deleteExpense mutation', async () => {
    const deleteButton = component.first().find('button[data-action="delete"]');
    expect(deleteButton).toExist();

    deleteButton.simulate('click');
    await updateComponent(component);
    expect(deleteMock.result).toHaveBeenCalled();
  });
});
