import React from 'react';

export const BudgetContext = React.createContext({});
export const useBudget = () => React.useContext(BudgetContext);
export { BudgetProvider } from './BudgetProvider';
