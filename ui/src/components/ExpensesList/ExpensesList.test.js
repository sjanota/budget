import {mount} from 'enzyme';
import React from 'react';
import {MockedProvider} from "@apollo/react-testing";
import {expense1, expense2, mockDeleteExpense, mockExpensesEvent, mockQueryExpenses} from "./ExpensesList.mocks";
import {updateComponent} from "../../testing";
import {createMockLink} from "../../testing/apollo";
import ExpensesList from "./ExpensesList";


it('displays loading initially', async () => {
  const {link} = createMockLink([]);
  console.error = jest.fn();
  const component = mount(
    <MockedProvider link={link}>
      <ExpensesList/>
    </MockedProvider>
  );

  expect(component.find('tbody tr')).toHaveLength(0);
  expect(component.find('p')).toExist();
  expect(component.find('p')).toHaveText("Loading...");
});

it('displays error if occurs', async () => {
  const {link} = createMockLink([]);
  console.error = jest.fn();
  const component = mount(
    <MockedProvider link={link}>
      <ExpensesList/>
    </MockedProvider>
  );
  await updateComponent(component);

  expect(console.error).toBeCalled();
  expect(component.find('tbody tr')).toHaveLength(0);
  expect(component.find('p')).toExist();
  expect(component.find('p').text).toMatchSnapshot();
});

it('displays queried data', async () => {
  const {link} = createMockLink([
    mockQueryExpenses([expense1]),
  ]);
  const component = mount(
    <MockedProvider link={link}>
      <ExpensesList/>
    </MockedProvider>
  );
  await updateComponent(component);

  expect(component.find('tbody tr')).toHaveLength(1);
});

it('updates list on ADDED', async () => {
  const {link, sendEvent} = createMockLink([
    mockQueryExpenses([expense1]),
  ]);
  const component = mount(
    <MockedProvider link={link}>
      <ExpensesList/>
    </MockedProvider>
  );
  sendEvent(mockExpensesEvent({type: 'ADDED', expense: expense2}));
  await updateComponent(component);
  await updateComponent(component);

  expect(component.find('tbody tr')).toHaveLength(2);
});

it('triggers deleteExpense mutation', async () => {
  const deleteMock = mockDeleteExpense(expense1.id);
  const {link} = createMockLink([
    mockQueryExpenses([expense1]),
    deleteMock
  ]);
  const component = mount(
    <MockedProvider link={link}>
      <ExpensesList/>
    </MockedProvider>
  );
  await updateComponent(component);

  const deleteButton = component.find('button[data-action="delete"]');
  expect(deleteButton).toExist();

  deleteButton.simulate('click');
  await updateComponent(component);
  expect(deleteMock.result).toHaveBeenCalled();
});

it('updates list on DELETED', async () => {
  const {link, sendEvent} = createMockLink([
    mockQueryExpenses([expense1, expense2]),
  ]);
  const component = mount(
    <MockedProvider link={link}>
      <ExpensesList/>
    </MockedProvider>
  );
  sendEvent(mockExpensesEvent({type: 'DELETED', expense: expense2}));
  await updateComponent(component);
  await updateComponent(component);

  expect(component.find('tbody tr')).toHaveLength(1);
});

