import {mount} from 'enzyme';
import App from "./App";
import React from 'react';
import {MockedProvider} from "@apollo/react-testing";
import {expense1, expense2, mockExpensesEvent, mockQueryExpenses} from "./App.mocks";
import {updateComponent} from "../../testing";
import {createMockLink} from "../../testing/apollo";


it('displays loading initially', async () => {
  const {link} = createMockLink([]);
  console.error = jest.fn();
  const component = mount(<App/>, {
    wrappingComponent: MockedProvider,
    wrappingComponentProps: {link}
  });

  expect(component.find('pre')).toHaveLength(0);
  expect(component.find('p')).toExist();
  expect(component.find('p')).toHaveText("Loading...");
});

it('displays error if occurs', async () => {
  const {link} = createMockLink([]);
  console.error = jest.fn();
  const component = mount(
    <MockedProvider link={link}>
      <App/>
    </MockedProvider>
  );
  await updateComponent(component);

  expect(console.error).toBeCalled();
  expect(component.find('pre')).toHaveLength(0);
  expect(component.find('p')).toExist();
  expect(component.find('p').text).toMatchSnapshot();
});

it('displays queried data', async () => {
  const {link} = createMockLink([
    mockQueryExpenses([expense1]),
  ]);
  const component = mount(
    <MockedProvider link={link}>
      <App/>
    </MockedProvider>
  );
  await updateComponent(component);

  expect(component.find('pre')).toHaveLength(1);
});

it('updates in ADDED', async () => {
  const {link, sendEvent} = createMockLink([
    mockQueryExpenses([expense1]),
  ]);
  const component = mount(
    <MockedProvider link={link}>
      <App/>
    </MockedProvider>
  );
  sendEvent(mockExpensesEvent({type: 'ADDED', expense: expense2, __typename: 'ExpenseAdded'}));
  await updateComponent(component);
  await updateComponent(component);

  expect(component.find('pre')).toHaveLength(2);
});



