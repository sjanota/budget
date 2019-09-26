import React from 'react';
import PropTypes from 'prop-types';
import { BudgetContext } from './budget';

export const MOCK_BUDGET_ID = 'fakebudegtid';

export const MockedBudgetProvider = ({ children }) => (
  <BudgetContext.Provider value={{ id: MOCK_BUDGET_ID }}>
    {children}
  </BudgetContext.Provider>
);

MockedBudgetProvider.propTypes = {
  children: PropTypes.any,
};
