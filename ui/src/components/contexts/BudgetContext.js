import React, { createContext, useState, useContext, useEffect } from 'react';
import PropTypes from 'prop-types';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

export const BudgetContext = createContext();
export const useBudget = () => useContext(BudgetContext);

const GET_BUDGETS = gql`
  query GetBudgets {
    budgets {
      id
      name
    }
  }
`;

export function BudgetProvider({ children }) {
  const [selectedBudget, setSelectedBudget] = useState(null);
  const { loading, error, data } = useQuery(GET_BUDGETS, {
    pollInterval: 10000,
  });
  const value = {
    selectedBudget,
    setSelectedBudget,
    loading,
    error,
    budgets: !loading && !error ? data.budgets : [],
  };
  if (error) {
    console.error(error);
  }
  return (
    <BudgetContext.Provider value={value}>{children}</BudgetContext.Provider>
  );
}

BudgetProvider.propTypes = {
  children: PropTypes.node,
};
